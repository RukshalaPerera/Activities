package handler

import (
	"Project3/app/db"
	"Project3/app/model"
	"errors"
	"github.com/gofiber/fiber/v2"
)

func GetReservationByID(id uint) (*model.Reservation, error) {
	var reservation model.Reservation
	if err := db.DB.First(&reservation, id).Error; err != nil {
		return nil, err // Return error if reservation not found
	}
	return &reservation, nil
}

func GetAllReservations(c *fiber.Ctx) error {
	var reservations []model.Reservation
	if err := db.DB.Find(&reservations).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(reservations)
}

func CreateReservation(c *fiber.Ctx) error {
	var newReservation model.Reservation
	if err := c.BodyParser(&newReservation); err != nil {
		return err
	}
	userID := newReservation.UserId
	var userReservations []model.Reservation
	if err := db.DB.Where("user_id = ?", userID).Find(&userReservations).Error; err != nil {
		return err
	}
	if len(userReservations) >= 4 {
		return fiber.NewError(fiber.StatusBadRequest, "You have reached the maximum limit of reservations")
	}

	var existingBook model.Book
	bookID := newReservation.BookId
	result := db.DB.First(&existingBook, bookID)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, "Book not found")
	}

	db.DB.Create(&newReservation)

	if existingBook.IsBookAvailable && newReservation.Completed() {
		if err := CompleteReservation(newReservation.Id); err != nil {
			return err
		}
	}

	return c.JSON(newReservation)
}

func UpdateReservation(c *fiber.Ctx) error {
	id := c.Params("id")
	var reservation model.Reservation
	if err := db.DB.First(&reservation, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Reservation not found"})
	}
	if err := c.BodyParser(&reservation); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := db.DB.Save(&reservation).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(reservation)
}

func DeleteReservation(c *fiber.Ctx) error {
	id := c.Params("id")
	var reservation model.Reservation
	if err := db.DB.First(&reservation, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Reservation not found"})
	}
	if err := db.DB.Delete(&reservation).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusOK)
}

func CompleteReservation(reservationID uint) error {
	reservation, err := GetReservationByID(reservationID)
	if err != nil {
		return err
	}
	if reservation == nil {
		return errors.New("reservation not found")
	}
	err = MarkReservationAsCompleted(reservationID)
	if err != nil {
		return err
	}
	return nil
}

func MarkReservationAsCompleted(reservationID uint) error {
	var reservation model.Reservation
	if err := db.DB.First(&reservation, reservationID).Error; err != nil {
		return err
	}
	reservation.Status = "Completed"

	if err := db.DB.Save(&reservation).Error; err != nil {
		return err
	}
	return nil
}
