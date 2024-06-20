package handler

import (
	"errors"
	"fmt"
	"github.com/bloomingbug/depublic/internal/entity"
	"github.com/bloomingbug/depublic/internal/http/binder"
	"github.com/bloomingbug/depublic/internal/http/form_validator"
	"github.com/bloomingbug/depublic/internal/service"
	"github.com/bloomingbug/depublic/internal/util"
	"github.com/bloomingbug/depublic/pkg/jwt_token"
	"github.com/bloomingbug/depublic/pkg/response"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type TransactionHandler struct {
	eventService       service.EventService
	timetableService   service.TimetableService
	transactionService service.TransactionService
	ticketService      service.TicketService
	paymentGateway     service.PaymenService
}

func (h *TransactionHandler) CreateTransaction(c echo.Context) error {
	req := new(binder.TransactionRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, response.Error(http.StatusBadRequest, false, form_validator.ValidatorErrors(err)))
	}

	event, err := h.eventService.FindEventById(c, req.EventID)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.Error(http.StatusUnprocessableEntity, false, err.Error()))
	}

	dataUser, _ := c.Get("user").(*jwt.Token)
	userClaims := dataUser.Claims.(*jwt_token.JwtCustomClaims)

	var ids []uuid.UUID
	ticketCounts := make(map[uuid.UUID]int32)
	for _, data := range req.Tickets {
		ids = append(ids, data.TimetableID)
		ticketCounts[data.TimetableID]++
	}

	timetables, err := h.timetableService.FindByIds(c, ids)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.Error(http.StatusUnprocessableEntity, false, err.Error()))
	}

	var amount int64 = 0
	if timetables != nil {
		for _, ticket := range req.Tickets {
			timetable, err := FilterTimetableByID(timetables, ticket.TimetableID)
			if err != nil {
				return c.JSON(http.StatusUnprocessableEntity, response.Error(http.StatusUnprocessableEntity, false, err.Error()))
			}
			var count int32 = ticketCounts[timetable.ID]

			if timetable.Stock < count {
				return c.JSON(http.StatusUnprocessableEntity, response.Error(http.StatusUnprocessableEntity, false, "Not enough stock"))
			}

			if err == nil {
				amount += *timetable.Price
			}
		}
	}

	userID, _ := uuid.Parse(userClaims.ID)
	var transactionParams = entity.NewTransactionParams{
		UserID:     userID,
		Invoice:    fmt.Sprintf("INVOICE-%s-%v", util.RandomStringGenerator(4), time.Now().Unix()),
		GrandTotal: amount,
	}

	transactionDTO := entity.NewTransaction(transactionParams)
	transaction, err := h.transactionService.CreateTransaction(c, transactionDTO)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.Error(http.StatusUnprocessableEntity, false, err.Error()))
	}

	var ticketsData []entity.Ticket
	for _, ticket := range req.Tickets {
		parsedDate, err := time.Parse("2006-01-02", ticket.Birthdate)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid birthdate format"})
		}

		timetable, err := FilterTimetableByID(timetables, ticket.TimetableID)
		if err != nil {
			return err
		}

		params := entity.NewTicketParams{
			Name:          ticket.Name,
			NoTicket:      fmt.Sprintf("DEPUBLIC-%s-%s", util.Abbreviate(event.Name), util.RandomStringGenerator(4)),
			PersonalNo:    ticket.PersonalNo,
			Birthdate:     parsedDate,
			Phone:         ticket.Phone,
			Email:         ticket.Email,
			Gender:        ticket.Gender,
			Price:         *timetable.Price,
			TimetableID:   ticket.TimetableID,
			TransactionID: transaction.ID,
		}
		ticketsData = append(ticketsData, *entity.NewTicket(params))
	}

	_, err = h.ticketService.CreateBatchTicket(c, transaction.ID, &ticketsData)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.Error(http.StatusUnprocessableEntity, false, err.Error()))
	}

	err = h.timetableService.UpdateTicketStock(c, ticketCounts, true)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Error(http.StatusInternalServerError, false, err.Error()))
	}

	paymentReq := entity.NewPayment(transaction.Invoice, transaction.GrandTotal, userClaims.Name, "", userClaims.Email)
	payment, err := h.paymentGateway.CreateTransaction(c.Request().Context(), paymentReq)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.Error(http.StatusUnprocessableEntity, false, err.Error()))
	}

	updateTransactionParams := entity.UpdateTransactionParams{
		ID:         transaction.ID,
		Invoice:    nil,
		GrandTotal: nil,
		SnapToken:  payment,
		Status:     nil,
	}

	updateTransaction := entity.UpdateTransaction(updateTransactionParams)
	_, err = h.transactionService.EditTransaction(c, updateTransaction)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Error(http.StatusInternalServerError, false, err.Error()))
	}

	return c.JSON(http.StatusCreated, response.Success(http.StatusCreated,
		true,
		"sukses menambahkan transaksi",
		echo.Map{"payment_url": payment}))
}

func FilterTimetableByID(data []entity.Timetable, id uuid.UUID) (*entity.Timetable, error) {
	for _, d := range data {
		if d.ID == id {
			return &d, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("data dengan id %d tidak ditemukan", id))
}

func NewTransactionHandler(eventService service.EventService,
	timetableService service.TimetableService,
	transactionService service.TransactionService,
	ticketService service.TicketService,
	paymentGateway service.PaymenService) TransactionHandler {
	return TransactionHandler{
		eventService:       eventService,
		timetableService:   timetableService,
		transactionService: transactionService,
		ticketService:      ticketService,
		paymentGateway:     paymentGateway,
	}
}
