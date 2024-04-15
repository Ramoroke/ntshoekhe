package drugs

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Drugs struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Quantity    int       `json:"quantity"`
	Location    string    `json:"location"`
	Expiry_Date time.Time `json:"expiry_date"`
}

var db *sql.DB

func Setup(database *sql.DB) {
	db = database
}

// func Listdrugss(c *gin.Context) {
// 	var drugss []drugs
// 	rows, err := db.Query("SELECT id, name, age, class FROM drugss")
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var u drugs
// 		if err := rows.Scan(&u.ID, &u.Name, &u.Age, &u.Class); err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}
// 		drugss = append(drugss, u)
// 	}
// 	c.JSON(http.StatusOK, drugss)
// }

func ListDrugs(c *gin.Context) {
	// Parse query parameters for pagination
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("perPage", "10"))

	if page < 1 {
		page = 1
	}

	offset := (page - 1) * perPage

	// Query to count total items
	var total int
	countQuery := "SELECT COUNT(*) FROM drugs"
	err := db.QueryRow(countQuery).Scan(&total)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Error retrieving resources.", "error": err.Error()})
		return
	}

	// Adjust totalPages to ensure it always has a value
	totalPages := total / perPage
	if total%perPage != 0 {
		totalPages++
	}

	// Query to fetch paginated items
	var drugs []Drugs
	query := "SELECT id, name, quantity, location FROM drugs LIMIT ? OFFSET ?"
	rows, err := db.Query(query, perPage, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Error retrieving resources.", "error": err.Error()})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var u Drugs
		if err := rows.Scan(&u.ID, &u.Name, &u.Quantity, &u.Location); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Error retrieving resources.", "error": err.Error()})
			return
		}
		//drugss = append(drugss, u)
		drugs = append(drugs, u)
	}

	// Construct the response
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Resources retrieved successfully.",
		"data": gin.H{
			"items":       drugs,
			"total":       total,
			"perPage":     perPage,
			"currentPage": page,
			"totalPages":  totalPages,
		},
	})
}

func AddDrug(c *gin.Context) {
	var newDrug Drugs
	if err := c.BindJSON(&newDrug); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Error creating resource.", "error": err.Error()})
		return
	}

	result, err := db.Exec("INSERT INTO drugs(name, qualtity, location, expiry_date) VALUES (?, ?, ?, ?)", newDrug.Name, newDrug.Quantity, newDrug.Location, newDrug.Expiry_Date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Error creating resource.", "error": err.Error()})
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Error creating resource.", "error": err.Error()})
		return
	}

	newDrug.ID = int(id)
	c.JSON(http.StatusCreated, gin.H{"success": true, "message": "Resource created successfully.", "data": newDrug})
}

func EditDrug(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid drugs ID.", "error": err.Error()})
		return
	}

	var updatedDrug Drugs
	if err := c.BindJSON(&updatedDrug); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Error updating resource.", "error": err.Error()})
		return
	}

	_, err = db.Exec("UPDATE drugs SET name = ?, quantity = ?, location = ?, expiry_data = ? WHERE id = ?", updatedDrug.Name, updatedDrug.Quantity, updatedDrug.Location, updatedDrug.Expiry_Date, id)
	if err != nil { // Assuming 'err' holds the error from your update operation
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Error updating resource.", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Resource updated successfully.", "data": updatedDrug})
}

func DeleteDrug(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Invalid drug ID.", "error": err.Error()})
		return
	}

	_, err = db.Exec("DELETE FROM drugs WHERE id = ?", id)
	if err != nil { // Assuming 'err' holds the error from your delete operation
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Error deleting resource.", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Resource deleted successfully."})
}

// CreateTable creates the drugss table if it does not exist
func CreateTable() {
	createTableSQL := `CREATE TABLE IF NOT EXISTS drugs (
        "id" INTEGER PRIMARY KEY AUTOINCREMENT,
        "name" TEXT NOT NULL,
        "quantity" INTEGER NOT NULL,
        "location" TEXT NOT NULL,
		"expiry_date" TEXT NOT NULL
    );`

	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Fatal("Failed to create table: ", err)
	}
}
