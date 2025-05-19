package handlers

import (
	"context"
	"crud_api/config"
	"crud_api/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetMahasiswa(c *fiber.Ctx) error {
	rows, err := config.DB.Query(context.Background(), "SELECT id, nama, nrp, jurusan_id, angkatan FROM mahasiswa")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "gagal ambil data"})
	}
	defer rows.Close()

	var mahasiswa []models.Mahasiswa
	for rows.Next() {
		var m models.Mahasiswa
		if err := rows.Scan(&m.ID, &m.Nama, &m.NRP, &m.JurusanID, &m.Angkatan); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "gagal scan data"})
		}

		mahasiswa = append(mahasiswa, m)
	}

	return c.Status(200).JSON(mahasiswa)
}

func CreateMahasiswa(c *fiber.Ctx) error {
	var m models.Mahasiswa

	if err := c.BodyParser(&m); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid input"})
	}

	query := `
		INSERT INTO mahasiswa (nama, nrp, jurusan_id, angkatan)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	err := config.DB.QueryRow(context.Background(), query, m.Nama, m.NRP, m.JurusanID, m.Angkatan).Scan(&m.ID)
	if err != nil {
		fmt.Println("Error inserting data:", err)
		return c.Status(500).JSON(fiber.Map{"error": "gagal simpan data"})

	}

	return c.Status(201).JSON(m)
}

func UpdateMahasiswa(c *fiber.Ctx) error {
	id := c.Params("id")
	var m models.Mahasiswa

	if err := c.BodyParser(&m); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid input"})
	}

	query := `
		UPDATE mahasiswa
		SET nama = $1, nrp = $2, jurusan_id = $3, angkatan = $4
		WHERE id = $5
	`

	_, err := config.DB.Exec(context.Background(), query, m.Nama, m.NRP, m.JurusanID, m.Angkatan, id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "gagal update data"})
	}

	return c.Status(200).JSON(m)
}

func DeleteMahasiswa(c *fiber.Ctx) error {
	id := c.Params("id")

	query := `
		DELETE FROM mahasiswa
		WHERE id = $1
	`

	_, err := config.DB.Exec(context.Background(), query, id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "gagal hapus data"})
	}

	return c.Status(200).JSON(fiber.Map{"message": "data berhasil dihapus"})
}
