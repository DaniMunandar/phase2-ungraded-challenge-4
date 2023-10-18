package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CriminalReportHandler struct {
	DB *sql.DB
}

type CriminalReport struct {
	ID          int     `json:"id"`
	Hero        Hero    `json:"hero"`
	Villain     Villain `json:"villain"`
	Description string  `json:"description"`
	Incident    string  `json:"incident"`
}

type Hero struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Universe string `json:"universe"`
	Skill    string `json:"skill"`
}

type Villain struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Universe string `json:"universe"`
}

func NewCriminalReportHandler(db *sql.DB) *CriminalReportHandler {
	return &CriminalReportHandler{DB: db}
}

func (h *CriminalReportHandler) GetCriminalReport(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	reportID, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	report, err := queryDatabaseToGetCriminalReportByID(h.DB, reportID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
	w.WriteHeader(http.StatusOK)
}

func (h *CriminalReportHandler) GetAllCriminalReports(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	reports, err := queryDatabaseToGetAllCriminalReports(h.DB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reports)
	w.WriteHeader(http.StatusOK)
}

func (h *CriminalReportHandler) CreateCriminalReport(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var report CriminalReport
	if err := json.NewDecoder(r.Body).Decode(&report); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdReport, err := saveCriminalReportToDatabase(h.DB, report)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdReport)
}

func (h *CriminalReportHandler) UpdateCriminalReport(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	reportID, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var updatedReport CriminalReport
	if err := json.NewDecoder(r.Body).Decode(&updatedReport); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedReport, err = updateCriminalReportInDatabase(h.DB, reportID, updatedReport)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedReport)
	w.WriteHeader(http.StatusOK)
}

func (h *CriminalReportHandler) DeleteCriminalReport(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	reportID, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = deleteCriminalReportFromDatabase(h.DB, reportID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func queryDatabaseToGetCriminalReportByID(db *sql.DB, reportID int) (CriminalReport, error) {
	var report CriminalReport
	err := db.QueryRow(`
        SELECT cr.id, cr.description, cr.incident, h.id, h.name, h.universe, h.skill, v.id, v.name, v.universe
        FROM CriminalReports cr
        JOIN Heroes h ON cr.hero_id = h.id
        JOIN Villain v ON cr.villain_id = v.id
        WHERE cr.id = ?`, reportID).Scan(&report.ID, &report.Description, &report.Incident, &report.Hero.ID, &report.Hero.Name, &report.Hero.Universe, &report.Hero.Skill, &report.Villain.ID, &report.Villain.Name, &report.Villain.Universe)
	if err != nil {
		return CriminalReport{}, err
	}
	return report, nil
}

func queryDatabaseToGetAllCriminalReports(db *sql.DB) ([]CriminalReport, error) {
	rows, err := db.Query(`
        SELECT cr.id, cr.description, cr.incident, h.id, h.name, h.universe, h.skill, v.id, v.name, v.universe
        FROM CriminalReports cr
        JOIN Heroes h ON cr.hero_id = h.id
        JOIN Villain v ON cr.villain_id = v.id
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reports []CriminalReport
	for rows.Next() {
		var report CriminalReport
		if err := rows.Scan(&report.ID, &report.Description, &report.Incident, &report.Hero.ID, &report.Hero.Name, &report.Hero.Universe, &report.Hero.Skill, &report.Villain.ID, &report.Villain.Name, &report.Villain.Universe); err != nil {
			return nil, err
		}
		reports = append(reports, report)
	}

	return reports, nil
}

func saveCriminalReportToDatabase(db *sql.DB, report CriminalReport) (CriminalReport, error) {
	result, err := db.Exec(`
        INSERT INTO CriminalReports (hero_id, villain_id, description, incident)
        VALUES (?, ?, ?, ?)`, report.Hero.ID, report.Villain.ID, report.Description, report.Incident)
	if err != nil {
		return CriminalReport{}, err
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return CriminalReport{}, err
	}
	report.ID = int(lastInsertID)
	return report, nil
}

func updateCriminalReportInDatabase(db *sql.DB, reportID int, updatedReport CriminalReport) (CriminalReport, error) {
	_, err := db.Exec(`
        UPDATE CriminalReports
        SET hero_id = ?, villain_id = ?, description = ?, incident = ?
        WHERE id = ?`, updatedReport.Hero.ID, updatedReport.Villain.ID, updatedReport.Description, updatedReport.Incident, reportID)
	if err != nil {
		return CriminalReport{}, err
	}
	updatedReport.ID = reportID
	return updatedReport, nil
}

func deleteCriminalReportFromDatabase(db *sql.DB, reportID int) error {
	_, err := db.Exec("DELETE FROM CriminalReports WHERE id = ?", reportID)
	return err
}
