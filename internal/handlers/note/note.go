package noteHandlers

import (
	"note-api-fiber/database"
	"note-api-fiber/internal/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetNotes(c *fiber.Ctx) error {
	db := database.DB
	var notes []models.Note

	db.Find(&notes)

	if len(notes) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "No notes found",
		})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "All notes", "data": notes})

}

func CreateNote(c *fiber.Ctx) error {
	db := database.DB
	note := new(models.Note)
	err := c.BodyParser(note)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Cannot parse JSON",
		})
	}
	note.ID = uuid.New()
	err = db.Create(&note).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Could not create note",
		})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created note", "data": note})
}

func GetNote(c *fiber.Ctx) error {
	db := database.DB
	var note models.Note

	db.Find(&note, "id = ?", c.Params("id"))

	if note.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "No note found with given ID",
		})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Note found", "data": note})
}

func UpdateNote(c *fiber.Ctx) error {
	db := database.DB
	var note models.Note

	db.Find(&note, "id = ?", c.Params("id"))

	if note.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "No note found with given ID",
		})
	}

	err := c.BodyParser(&note)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Cannot parse JSON",
		})
	}

	db.Save(&note)

	return c.JSON(fiber.Map{"status": "success", "message": "Note updated", "data": note})
}

func DeleteNote(c *fiber.Ctx) error {
	db := database.DB
	var note models.Note

	db.Find(&note, "id = ?", c.Params("id"))

	if note.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "No note found with given ID",
		})
	}

	db.Delete(&note)

	return c.JSON(fiber.Map{"status": "success", "message": "Note deleted"})

}
