package controller

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/internship1/wsMail/config"
	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	iniconfig "github.com/internship1/backendmail/config"
	model "github.com/internship1/backendmail/model"
	cek "github.com/internship1/backendmail/module"
)

func Homepage(c *fiber.Ctx) error {
	ipaddr := musik.GetIPaddress()
	return c.JSON(ipaddr)
}

func GenerateConnote() string {
	timestamp := time.Now().Format("020106") // format: ddmmyy
	randomNum := rand.Intn(9999999) //random 7 digit
	return fmt.Sprintf("P%s%07d", timestamp, randomNum)
}

// InsertDataTransaction godoc
// @Summary Insert Data Transaction.
// @Description Input data transaction.
// @Tags Transaction
// @Accept json
// @Produce json
// @Param request body ReqTransaction true "Payload Body [RAW]"
// @Success 200 {object} Transaction
// @Failure 400
// @Failure 500
// @Router /insertTransaction [post]
func InsertDataTransaction(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var transaction model.Transaction
	if err := c.BodyParser(&transaction); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	insertedID, err := cek.InsertTransaction(db, "MailApp",
		transaction.ConsignmentNote,
		transaction.SenderName,
		transaction.ReceiverName,
		transaction.AddressReceiver,
		transaction.ReceiverPhone,
		transaction.ItemContent,
		transaction.DeliveryStatus,
		transaction.CODValue,
		transaction.CreatedAt,
		transaction.UpdatedAt)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

// GetAllTransaction godoc
// @Summary Get All Data Transaction.
// @Description Mengambil semua data transaksi.
// @Tags MenuItem
// @Accept json
// @Produce json
// @Success 200 {object} Transaction
// @Router /transaction [get]
// GetAllTransaction retrieves all transaction from the database
// GetAllTransaction retrieves all transaction from the database
func GetAllTransaction(c *fiber.Ctx) error {
	ps := cek.GetAllTransaction(config.Ulbimongoconn, "MailApp")
	return c.JSON(ps)
}

// GetTransactionByConnote retrieves transactions by consignment note
// GetMenuID godoc
// @Summary Get By ID Data Menu.
// @Description Ambil per ID data menu.
// @Tags MenuItem
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200 {object} MenuItem
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /menu/{id} [get]
func GetTransactionByConnote(c *fiber.Ctx) error {
	connote := c.Params("connote")
	if connote == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Parameter connote tidak boleh kosong",
		})
	}

	transaction, err := cek.GetTransactionByConnote(config.Ulbimongoconn, "MailApp", connote)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"status":  http.StatusNotFound,
				"message": fmt.Sprintf("Tidak ada data dengan connote %s", connote),
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Gagal mengambil data untuk connote %s: %v", connote, err),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Berhasil mengambil data transaksi",
		"data":    transaction,
	})
}

// GetTransactionByPhoneNumber retrieves transactions by phone number
// GetTransactionPhoneNumber godoc
// @Summary Get By Phone Number Data Transaction.
// @Description Ambil per Nomor telepon data transaksi.
// @Tags Transaction
// @Accept json
// @Produce json
// @Param phoneNumber path string true "Masukan nomor telepon"
// @Success 200 {object} Transaction
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /tarnsaction/{phoneNumber} [get]
func GetTransactionByPhoneNumber(c *fiber.Ctx) error {
	phoneNumber := c.Params("phoneNumber")
	if phoneNumber == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Parameter connote tidak boleh kosong",
		})
	}

	transaction, err := cek.GetTransactionByPhoneNumber(config.Ulbimongoconn, "MailApp", phoneNumber)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"status":  http.StatusNotFound,
				"message": fmt.Sprintf("Tidak ada data dengan nomor telepon %s", phoneNumber),
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Gagal mengambil data untuk nomor telepon %s: %v", phoneNumber, err),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Berhasil mengambil data transaksi",
		"data":    transaction,
	})
}

// GetTransactionByAddress retrieves transactions by address
// GetTransactionAddress godoc
// @Summary Get By Address Data Transaction.
// @Description Ambil per alamat data transaksi.
// @Tags Transaction
// @Accept json
// @Produce json
// @Param address path string true "Masukan alamat"
// @Success 200 {object} Transaction
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /transaction/{address} [get]
func GetTransactionByAddress(c *fiber.Ctx) error {
	address := c.Params("address")
	if address == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Parameter address tidak boleh kosong",
		})
	}

	transaction, err := cek.GetTransactionByConnote(config.Ulbimongoconn, "MailApp", address)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"status":  http.StatusNotFound,
				"message": fmt.Sprintf("Tidak ada data dengan alamat %s", address),
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Gagal mengambil data untuk alamat %s: %v", address, err),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Berhasil mengambil data transaksi",
		"data":    transaction,
	})
}

// GetUser godoc
// @Summary Get All Data User.
// @Description Mengambil semua data user.
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} User
// @Router /user [get]
func GetAllUser(c *fiber.Ctx) error {
	ps, err := cek.GetAllUser(config.Ulbimongoconn, "User")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.JSON(ps)
}

// GetUserID godoc
// @Summary Get By ID Data User.
// @Description Ambil per ID data USER.
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200 {object} User
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /user/{id} [get]
func GetUserID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}
	ps, err := cek.GetUserByID(objID, config.Ulbimongoconn, "User")
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"status":  http.StatusNotFound,
				"message": fmt.Sprintf("No data found for id %s", id),
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error retrieving data for id %s", id),
		})
	}
	return c.JSON(ps)
}

// InsertDataUser godoc
// @Summary Insert data user.
// @Description Input data user.
// @Tags User
// @Accept json
// @Produce json
// @Param request body ReqUser true "Payload Body [RAW]"
// @Success 200 {object} User
// @Failure 400
// @Failure 500
// @Router /insertUser [post]
func InsertDataUser(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var user model.User

	// Parsing body request ke struct User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
	}

	// Hash password sebelum disimpan
	hashedPassword, err := iniconfig.HashPassword(user.Password)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Could not hash password",
		})
	}
	user.Password = hashedPassword

	// Pastikan PersonalizedCategories tidak nil
	if user.PersonalizedCategories == nil {
		user.PersonalizedCategories = []string{}
	}

	// Insert ke database
	insertedID, err := cek.InsertUser(db, "User",
		user.FullName,
		user.Phone,
		user.Username,
		user.Password,
		user.Role,
		user.PersonalizedCategories, // Tidak perlu [] di sini
	)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Response sukses
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

// UpdateData godoc
// @Summary Update data user.
// @Description Ubah data user.
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Param request body ReqUser true "Payload Body [RAW]"
// @Success 200 {object} User
// @Failure 400
// @Failure 500
// @Router /updateUser/{id} [put]
func UpdateDataUser(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	id := c.Params("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Hash password sebelum disimpan
	hashedPassword, err := iniconfig.HashPassword(user.Password)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Could not hash password",
		})
	}
	user.Password = hashedPassword

	err = cek.UpdateUser(context.Background(), db, "User",
		objectID,
		user.FullName,
		user.Phone,
		user.Username,
		user.Password,
		user.Role)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Data successfully updated",
	})
}

// DeleteUserByID godoc
// @Summary Delete data user.
// @Description Hapus data user.
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "Masukan ID"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /deleteUser/{id} [delete]
func DeleteUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}

	err = cek.DeleteUserByID(objID, config.Ulbimongoconn, "User")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error deleting data for id %s", id),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": fmt.Sprintf("Data with id %s deleted successfully", id),
	})
}
