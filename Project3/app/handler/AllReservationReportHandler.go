package handler

import (
	"Project3/app/db"
	"Project3/app/model"
	"github.com/gofiber/fiber/v2"
	"github.com/jung-kurt/gofpdf"
	"strconv"
	"time"
)

func GenerateAllReservationReportHandler(c *fiber.Ctx) error {
	startTimeStr := c.Query("start_time")
	endTimeStr := c.Query("end_time")

	var startTime, endTime time.Time
	var err error
	if startTimeStr != "" {
		startTime, err = time.Parse(time.RFC822Z, startTimeStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid start_time format"})
		}
	}
	if endTimeStr != "" {
		endTime, err = time.Parse(time.RFC822Z, endTimeStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid end_time format"})
		}
	}

	query := db.DB.Model(&model.Reservation{})
	if !startTime.IsZero() {
		query = query.Where("start_time >= ?", startTime)
	}
	if !endTime.IsZero() {
		query = query.Where("end_time <= ?", endTime)
	}

	var reservations []model.Reservation
	if err := query.Find(&reservations).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(40, 10, "All Reservation Report\n\n")

	if !startTime.IsZero() || !endTime.IsZero() {
		pdf.Cell(0, 10, "Time Range:\n\n")
		if !startTime.IsZero() {
			pdf.Cell(0, 10, "Start: "+startTime.Format(time.RFC3339)+"\n\n")
		}
		if !endTime.IsZero() {
			pdf.Cell(0, 10, "End: "+endTime.Format(time.RFC3339)+"\n\n")
		}
	}
	generatedAt := time.Now().Format("2006-01-02 15:04:05")
	for _, reservation := range reservations {
		pdf.Cell(0, 10, "Book ID: "+strconv.Itoa(int(reservation.BookId))+"\n\n")
		pdf.Cell(0, 10, "User ID: "+strconv.Itoa(int(reservation.UserId))+"\n\n")
		pdf.Cell(0, 10, "Is Complete: "+strconv.FormatBool(reservation.IsCompleted)+"\n\n")
		pdf.Cell(0, 10, "Reservation Start Time: "+reservation.StartTime.Format(time.RFC3339)+"\n\n")
		pdf.Cell(0, 10, "Reservation End Time: "+reservation.EndTime.Format(time.RFC3339)+"\n\n")
		pdf.Cell(0, 10, "Status: "+reservation.Status+"\n\n")
	}
	pdf.Cell(0, 10, "Generated At: "+generatedAt+"\n\n")

	if err := pdf.OutputFileAndClose("BookReport.pdf"); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}
	return c.SendFile("BookReport.pdf")
}
