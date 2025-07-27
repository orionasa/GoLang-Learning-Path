# 🚀 Module 8.2: CSV Processing

<div align="center">

![Go Logo](https://golang.org/lib/godoc/images/go-logo-blue.svg)

**Master CSV Data Processing in Go**

*From Zero to Go Hero - Module 8.2*

</div>

---

## 📋 Lesson Overview

| **Duration** | **Difficulty** | **Prerequisites** |
|--------------|----------------|-------------------|
| 1-1.5 hours  | Intermediate   | File I/O, Structs, Error Handling |

### 🎯 Learning Objectives
By the end of this lesson, you will:
- ✅ Read and write CSV files efficiently in Go
- ✅ Handle CSV headers and custom field mappings
- ✅ Process large CSV files with streaming techniques
- ✅ Validate and transform CSV data
- ✅ Build a complete data analysis pipeline

---

## 🌟 What is CSV?

<div style="background: linear-gradient(135deg, #f093fb 0%, #f5576c 100%); padding: 20px; border-radius: 10px; color: white; margin: 20px 0;">

### 📊 **CSV (Comma-Separated Values)**

CSV is a simple file format used to store tabular data, like a spreadsheet or database table. Each line represents a row, and columns are separated by commas (or other delimiters).

**Think of CSV like a digital spreadsheet** - it's the most common way to exchange data between different systems!

</div>

### 🏗️ **CSV Structure Example**
```
┌─────────────────────────────────────┐
│            CSV Structure            │
├─────────────────────────────────────┤
│  Header Row:  name,age,city,salary  │
│  Data Row 1:  John,25,NYC,50000     │
│  Data Row 2:  Alice,30,LA,60000     │
│  Data Row 3:  Bob,28,Chicago,55000  │
└─────────────────────────────────────┘
```

---

## 🔧 CSV in Go: The `encoding/csv` Package

<div style="background: #f8f9fa; border-left: 4px solid #28a745; padding: 15px; margin: 10px 0;">

**Key Components** - Go's CSV package provides:

```go
// Reader for parsing CSV data
type Reader struct { ... }

// Writer for creating CSV data  
type Writer struct { ... }

// Main methods
func (r *Reader) Read() ([]string, error)
func (r *Reader) ReadAll() ([][]string, error)
func (w *Writer) Write(record []string) error
```

</div>

### 🎯 **Basic CSV Reading**

Let's start with reading a simple CSV file:

```go
package main

import (
    "encoding/csv"
    "fmt"
    "io"
    "log"
    "strings"
)

func main() {
    // Sample CSV data
    csvData := `name,age,city,salary
John Doe,25,New York,50000
Alice Smith,30,Los Angeles,60000
Bob Johnson,28,Chicago,55000
Diana Brown,32,Houston,65000`

    // Create a CSV reader
    reader := csv.NewReader(strings.NewReader(csvData))
    
    // Read header
    header, err := reader.Read()
    if err != nil {
        log.Fatal("Error reading header:", err)
    }
    fmt.Println("Header:", header)
    fmt.Println("=" * 40)

    // Read all records
    records, err := reader.ReadAll()
    if err != nil {
        log.Fatal("Error reading records:", err)
    }

    // Display data in a formatted way
    for i, record := range records {
        fmt.Printf("Employee %d:\n", i+1)
        for j, field := range record {
            fmt.Printf("  %s: %s\n", header[j], field)
        }
        fmt.Println()
    }
}
```

**Output:**
```
Header: [name age city salary]
========================================
Employee 1:
  name: John Doe
  age: 25
  city: New York
  salary: 50000

Employee 2:
  name: Alice Smith
  age: 30
  city: Los Angeles
  salary: 60000
...
```

---

## 🏗️ Structured CSV Processing with Structs

<div style="background: linear-gradient(135deg, #4facfe 0%, #00f2fe 100%); padding: 20px; border-radius: 10px; color: white; margin: 20px 0;">

### 🎨 **From CSV to Go Structs**

Instead of working with raw string slices, let's convert CSV data into proper Go structs for type safety and easier manipulation.

</div>

### 🔧 **Employee Management System**

```go
package main

import (
    "encoding/csv"
    "fmt"
    "io"
    "log"
    "os"
    "strconv"
    "strings"
    "time"
)

// Employee represents an employee record
type Employee struct {
    ID          int       `csv:"id"`
    Name        string    `csv:"name"`
    Age         int       `csv:"age"`
    Department  string    `csv:"department"`
    Salary      float64   `csv:"salary"`
    HireDate    time.Time `csv:"hire_date"`
    IsActive    bool      `csv:"is_active"`
}

// EmployeeProcessor handles CSV operations for employees
type EmployeeProcessor struct {
    employees []Employee
}

// NewEmployeeProcessor creates a new processor
func NewEmployeeProcessor() *EmployeeProcessor {
    return &EmployeeProcessor{
        employees: make([]Employee, 0),
    }
}

// LoadFromCSV reads employees from a CSV file
func (ep *EmployeeProcessor) LoadFromCSV(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return fmt.Errorf("failed to open file: %w", err)
    }
    defer file.Close()

    reader := csv.NewReader(file)
    
    // Read header
    header, err := reader.Read()
    if err != nil {
        return fmt.Errorf("failed to read header: %w", err)
    }

    // Create field index map
    fieldIndex := make(map[string]int)
    for i, field := range header {
        fieldIndex[field] = i
    }

    // Read and parse records
    for {
        record, err := reader.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            return fmt.Errorf("failed to read record: %w", err)
        }

        employee, err := ep.parseEmployee(record, fieldIndex)
        if err != nil {
            log.Printf("Warning: skipping invalid record: %v", err)
            continue
        }

        ep.employees = append(ep.employees, employee)
    }

    return nil
}

// parseEmployee converts a CSV record to an Employee struct
func (ep *EmployeeProcessor) parseEmployee(record []string, fieldIndex map[string]int) (Employee, error) {
    var emp Employee
    var err error

    // Parse ID
    if idx, exists := fieldIndex["id"]; exists && idx < len(record) {
        emp.ID, err = strconv.Atoi(record[idx])
        if err != nil {
            return emp, fmt.Errorf("invalid ID: %w", err)
        }
    }

    // Parse Name
    if idx, exists := fieldIndex["name"]; exists && idx < len(record) {
        emp.Name = strings.TrimSpace(record[idx])
    }

    // Parse Age
    if idx, exists := fieldIndex["age"]; exists && idx < len(record) {
        emp.Age, err = strconv.Atoi(record[idx])
        if err != nil {
            return emp, fmt.Errorf("invalid age: %w", err)
        }
    }

    // Parse Department
    if idx, exists := fieldIndex["department"]; exists && idx < len(record) {
        emp.Department = strings.TrimSpace(record[idx])
    }

    // Parse Salary
    if idx, exists := fieldIndex["salary"]; exists && idx < len(record) {
        emp.Salary, err = strconv.ParseFloat(record[idx], 64)
        if err != nil {
            return emp, fmt.Errorf("invalid salary: %w", err)
        }
    }

    // Parse Hire Date
    if idx, exists := fieldIndex["hire_date"]; exists && idx < len(record) {
        emp.HireDate, err = time.Parse("2006-01-02", record[idx])
        if err != nil {
            return emp, fmt.Errorf("invalid hire date: %w", err)
        }
    }

    // Parse IsActive
    if idx, exists := fieldIndex["is_active"]; exists && idx < len(record) {
        emp.IsActive, err = strconv.ParseBool(record[idx])
        if err != nil {
            return emp, fmt.Errorf("invalid is_active: %w", err)
        }
    }

    return emp, nil
}

// SaveToCSV writes employees to a CSV file
func (ep *EmployeeProcessor) SaveToCSV(filename string) error {
    file, err := os.Create(filename)
    if err != nil {
        return fmt.Errorf("failed to create file: %w", err)
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    // Write header
    header := []string{"id", "name", "age", "department", "salary", "hire_date", "is_active"}
    if err := writer.Write(header); err != nil {
        return fmt.Errorf("failed to write header: %w", err)
    }

    // Write records
    for _, emp := range ep.employees {
        record := []string{
            strconv.Itoa(emp.ID),
            emp.Name,
            strconv.Itoa(emp.Age),
            emp.Department,
            fmt.Sprintf("%.2f", emp.Salary),
            emp.HireDate.Format("2006-01-02"),
            strconv.FormatBool(emp.IsActive),
        }

        if err := writer.Write(record); err != nil {
            return fmt.Errorf("failed to write record: %w", err)
        }
    }

    return nil
}

// GetEmployees returns all employees
func (ep *EmployeeProcessor) GetEmployees() []Employee {
    return ep.employees
}

// FilterByDepartment returns employees from a specific department
func (ep *EmployeeProcessor) FilterByDepartment(department string) []Employee {
    var filtered []Employee
    for _, emp := range ep.employees {
        if strings.EqualFold(emp.Department, department) {
            filtered = append(filtered, emp)
        }
    }
    return filtered
}

// CalculateAverageSalary calculates average salary by department
func (ep *EmployeeProcessor) CalculateAverageSalary() map[string]float64 {
    departmentSums := make(map[string]float64)
    departmentCounts := make(map[string]int)

    for _, emp := range ep.employees {
        if emp.IsActive {
            departmentSums[emp.Department] += emp.Salary
            departmentCounts[emp.Department]++
        }
    }

    averages := make(map[string]float64)
    for dept, sum := range departmentSums {
        averages[dept] = sum / float64(departmentCounts[dept])
    }

    return averages
}

// DisplaySummary shows a summary of employee data
func (ep *EmployeeProcessor) DisplaySummary() {
    fmt.Println("👥 EMPLOYEE SUMMARY")
    fmt.Println("==================")
    fmt.Printf("Total Employees: %d\n", len(ep.employees))
    
    activeCount := 0
    for _, emp := range ep.employees {
        if emp.IsActive {
            activeCount++
        }
    }
    fmt.Printf("Active Employees: %d\n", activeCount)
    fmt.Printf("Inactive Employees: %d\n", len(ep.employees)-activeCount)
    fmt.Println()

    // Department breakdown
    deptCounts := make(map[string]int)
    for _, emp := range ep.employees {
        if emp.IsActive {
            deptCounts[emp.Department]++
        }
    }

    fmt.Println("📊 DEPARTMENT BREAKDOWN")
    for dept, count := range deptCounts {
        fmt.Printf("  %s: %d employees\n", dept, count)
    }
    fmt.Println()

    // Average salaries
    avgSalaries := ep.CalculateAverageSalary()
    fmt.Println("💰 AVERAGE SALARIES BY DEPARTMENT")
    for dept, avg := range avgSalaries {
        fmt.Printf("  %s: $%.2f\n", dept, avg)
    }
}

func main() {
    // Create sample CSV data
    csvData := `id,name,age,department,salary,hire_date,is_active
1,John Doe,25,Engineering,75000.00,2022-01-15,true
2,Alice Smith,30,Marketing,65000.00,2021-03-20,true
3,Bob Johnson,28,Engineering,80000.00,2021-11-10,true
4,Diana Brown,32,Sales,70000.00,2020-05-12,true
5,Charlie Wilson,29,Marketing,62000.00,2022-08-01,false
6,Eva Davis,27,Engineering,78000.00,2023-02-14,true`

    // Write sample data to file
    file, err := os.Create("employees.csv")
    if err != nil {
        log.Fatal("Error creating file:", err)
    }
    file.WriteString(csvData)
    file.Close()

    // Process the CSV file
    processor := NewEmployeeProcessor()
    
    if err := processor.LoadFromCSV("employees.csv"); err != nil {
        log.Fatal("Error loading CSV:", err)
    }

    // Display summary
    processor.DisplaySummary()

    // Filter by department
    fmt.Println("\n🔧 ENGINEERING TEAM")
    fmt.Println("===================")
    engineers := processor.FilterByDepartment("Engineering")
    for _, emp := range engineers {
        fmt.Printf("  %s (Age: %d, Salary: $%.2f)\n", 
            emp.Name, emp.Age, emp.Salary)
    }

    // Save filtered data
    engineerProcessor := NewEmployeeProcessor()
    engineerProcessor.employees = engineers
    
    if err := engineerProcessor.SaveToCSV("engineers.csv"); err != nil {
        log.Fatal("Error saving filtered CSV:", err)
    }
    
    fmt.Println("\n✅ Filtered data saved to engineers.csv")

    // Clean up
    os.Remove("employees.csv")
    os.Remove("engineers.csv")
}
```

---

## 🌊 Streaming Large CSV Files

<div style="background: linear-gradient(135deg, #fa709a 0%, #fee140 100%); padding: 20px; border-radius: 10px; color: white; margin: 20px 0;">

### 🚀 **Memory-Efficient Processing**

When dealing with large CSV files (millions of rows), loading everything into memory isn't practical. Let's use streaming techniques to process data efficiently.

</div>

### 🔧 **Large File Processor**

```go
package main

import (
    "encoding/csv"
    "fmt"
    "io"
    "log"
    "os"
    "strconv"
    "strings"
)

// SalesRecord represents a sales transaction
type SalesRecord struct {
    ID         int     `csv:"id"`
    Date       string  `csv:"date"`
    Product    string  `csv:"product"`
    Category   string  `csv:"category"`
    Amount     float64 `csv:"amount"`
    Quantity   int     `csv:"quantity"`
    CustomerID int     `csv:"customer_id"`
}

// SalesAnalyzer processes sales data in streaming fashion
type SalesAnalyzer struct {
    totalSales     float64
    recordCount    int
    categoryTotals map[string]float64
    productCounts  map[string]int
}

// NewSalesAnalyzer creates a new analyzer
func NewSalesAnalyzer() *SalesAnalyzer {
    return &SalesAnalyzer{
        categoryTotals: make(map[string]float64),
        productCounts:  make(map[string]int),
    }
}

// ProcessCSVStream processes CSV file in streaming mode
func (sa *SalesAnalyzer) ProcessCSVStream(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return fmt.Errorf("failed to open file: %w", err)
    }
    defer file.Close()

    reader := csv.NewReader(file)
    
    // Read and skip header
    header, err := reader.Read()
    if err != nil {
        return fmt.Errorf("failed to read header: %w", err)
    }
    
    fmt.Printf("📊 Processing CSV with columns: %v\n", header)
    fmt.Println("🔄 Streaming data...")

    // Create field index map
    fieldIndex := make(map[string]int)
    for i, field := range header {
        fieldIndex[field] = i
    }

    // Process records one by one
    batchSize := 1000
    processed := 0
    
    for {
        record, err := reader.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            log.Printf("Warning: skipping invalid record at line %d: %v", 
                processed+2, err) // +2 for header and 1-based indexing
            continue
        }

        // Parse and process record
        salesRecord, err := sa.parseRecord(record, fieldIndex)
        if err != nil {
            log.Printf("Warning: skipping invalid record at line %d: %v", 
                processed+2, err)
            continue
        }

        sa.processRecord(salesRecord)
        processed++

        // Show progress for large files
        if processed%batchSize == 0 {
            fmt.Printf("  ✅ Processed %d records...\n", processed)
        }
    }

    fmt.Printf("🎉 Completed! Processed %d total records\n\n", processed)
    return nil
}

// parseRecord converts CSV record to SalesRecord
func (sa *SalesAnalyzer) parseRecord(record []string, fieldIndex map[string]int) (SalesRecord, error) {
    var sr SalesRecord
    var err error

    // Parse ID
    if idx, exists := fieldIndex["id"]; exists && idx < len(record) {
        sr.ID, err = strconv.Atoi(record[idx])
        if err != nil {
            return sr, fmt.Errorf("invalid ID: %w", err)
        }
    }

    // Parse Date
    if idx, exists := fieldIndex["date"]; exists && idx < len(record) {
        sr.Date = strings.TrimSpace(record[idx])
    }

    // Parse Product
    if idx, exists := fieldIndex["product"]; exists && idx < len(record) {
        sr.Product = strings.TrimSpace(record[idx])
    }

    // Parse Category
    if idx, exists := fieldIndex["category"]; exists && idx < len(record) {
        sr.Category = strings.TrimSpace(record[idx])
    }

    // Parse Amount
    if idx, exists := fieldIndex["amount"]; exists && idx < len(record) {
        sr.Amount, err = strconv.ParseFloat(record[idx], 64)
        if err != nil {
            return sr, fmt.Errorf("invalid amount: %w", err)
        }
    }

    // Parse Quantity
    if idx, exists := fieldIndex["quantity"]; exists && idx < len(record) {
        sr.Quantity, err = strconv.Atoi(record[idx])
        if err != nil {
            return sr, fmt.Errorf("invalid quantity: %w", err)
        }
    }

    // Parse Customer ID
    if idx, exists := fieldIndex["customer_id"]; exists && idx < len(record) {
        sr.CustomerID, err = strconv.Atoi(record[idx])
        if err != nil {
            return sr, fmt.Errorf("invalid customer_id: %w", err)
        }
    }

    return sr, nil
}

// processRecord updates analytics with the record data
func (sa *SalesAnalyzer) processRecord(record SalesRecord) {
    sa.totalSales += record.Amount
    sa.recordCount++
    sa.categoryTotals[record.Category] += record.Amount
    sa.productCounts[record.Product] += record.Quantity
}

// GenerateReport creates a summary report
func (sa *SalesAnalyzer) GenerateReport() {
    fmt.Println("📈 SALES ANALYSIS REPORT")
    fmt.Println("========================")
    fmt.Printf("Total Records Processed: %d\n", sa.recordCount)
    fmt.Printf("Total Sales Amount: $%.2f\n", sa.totalSales)
    fmt.Printf("Average Sale Amount: $%.2f\n", sa.totalSales/float64(sa.recordCount))
    fmt.Println()

    fmt.Println("💰 SALES BY CATEGORY")
    fmt.Println("--------------------")
    for category, total := range sa.categoryTotals {
        percentage := (total / sa.totalSales) * 100
        fmt.Printf("  %s: $%.2f (%.1f%%)\n", category, total, percentage)
    }
    fmt.Println()

    fmt.Println("📦 TOP PRODUCTS BY QUANTITY")
    fmt.Println("---------------------------")
    // Find top 5 products
    type ProductCount struct {
        Product string
        Count   int
    }
    
    var products []ProductCount
    for product, count := range sa.productCounts {
        products = append(products, ProductCount{product, count})
    }
    
    // Simple bubble sort for top products
    for i := 0; i < len(products)-1; i++ {
        for j := 0; j < len(products)-i-1; j++ {
            if products[j].Count < products[j+1].Count {
                products[j], products[j+1] = products[j+1], products[j]
            }
        }
    }
    
    // Show top 5
    limit := 5
    if len(products) < limit {
        limit = len(products)
    }
    
    for i := 0; i < limit; i++ {
        fmt.Printf("  %d. %s: %d units\n", i+1, products[i].Product, products[i].Count)
    }
}

// ExportSummaryCSV exports analysis results to CSV
func (sa *SalesAnalyzer) ExportSummaryCSV(filename string) error {
    file, err := os.Create(filename)
    if err != nil {
        return fmt.Errorf("failed to create file: %w", err)
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    // Write category summary
    writer.Write([]string{"Category", "Total_Sales", "Percentage"})
    for category, total := range sa.categoryTotals {
        percentage := (total / sa.totalSales) * 100
        record := []string{
            category,
            fmt.Sprintf("%.2f", total),
            fmt.Sprintf("%.1f", percentage),
        }
        writer.Write(record)
    }

    return nil
}

func main() {
    // Create sample large CSV data
    csvData := `id,date,product,category,amount,quantity,customer_id
1,2024-01-01,Laptop,Electronics,1200.00,1,101
2,2024-01-01,Coffee Mug,Home,15.99,2,102
3,2024-01-02,Smartphone,Electronics,899.99,1,103
4,2024-01-02,Book,Education,29.99,1,104
5,2024-01-03,Headphones,Electronics,199.99,1,105
6,2024-01-03,Desk Chair,Furniture,299.99,1,106
7,2024-01-04,Notebook,Education,12.99,3,107
8,2024-01-04,Monitor,Electronics,449.99,1,108
9,2024-01-05,Coffee Beans,Food,24.99,2,109
10,2024-01-05,Keyboard,Electronics,89.99,1,110`

    // Write sample data to file
    file, err := os.Create("sales_data.csv")
    if err != nil {
        log.Fatal("Error creating file:", err)
    }
    file.WriteString(csvData)
    file.Close()

    // Process the data
    analyzer := NewSalesAnalyzer()
    
    if err := analyzer.ProcessCSVStream("sales_data.csv"); err != nil {
        log.Fatal("Error processing CSV:", err)
    }

    // Generate report
    analyzer.GenerateReport()

    // Export summary
    if err := analyzer.ExportSummaryCSV("sales_summary.csv"); err != nil {
        log.Fatal("Error exporting summary:", err)
    }
    
    fmt.Println("\n✅ Summary exported to sales_summary.csv")

    // Clean up
    os.Remove("sales_data.csv")
    os.Remove("sales_summary.csv")
}
```

---

## 🛠️ Advanced CSV Techniques

### 🎯 **Custom Delimiters and Formats**

```go
package main

import (
    "encoding/csv"
    "fmt"
    "strings"
)

func main() {
    // CSV with different delimiter (semicolon)
    csvData := `name;age;city;country
John Doe;25;New York;USA
Alice Smith;30;London;UK
Bob Johnson;28;Tokyo;Japan`

    reader := csv.NewReader(strings.NewReader(csvData))
    reader.Comma = ';' // Set custom delimiter
    reader.Comment = '#' // Lines starting with # are comments
    reader.TrimLeadingSpace = true // Remove leading spaces

    records, err := reader.ReadAll()
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }

    fmt.Println("📋 PARSED DATA (Semicolon Delimiter)")
    fmt.Println("====================================")
    for i, record := range records {
        if i == 0 {
            fmt.Printf("Headers: %v\n", record)
            continue
        }
        fmt.Printf("Record %d: %v\n", i, record)
    }
}
```

---

## 🎯 Practice Exercises

### 🏋️ **Exercise 1: Student Grade Processor**
Create a system that:
1. Reads student grades from CSV
2. Calculates GPA for each student
3. Generates class statistics
4. Exports results in multiple formats

### 🏋️ **Exercise 2: Financial Data Analyzer**
Build a financial analyzer that:
1. Processes transaction CSV files
2. Categorizes expenses and income
3. Generates monthly/yearly reports
4. Handles different CSV formats and currencies

### 🏋️ **Exercise 3: Log File Processor**
Create a log processor that:
1. Parses CSV log files with timestamps
2. Filters by date ranges and severity levels
3. Generates summary statistics
4. Exports filtered results

---

## 🎉 Summary

Congratulations! You've mastered CSV processing in Go. Here's what you've learned:

<div style="background: linear-gradient(135deg, #f093fb 0%, #f5576c 100%); padding: 20px; border-radius: 10px; color: white; margin: 20px 0;">

### 🎯 **Key Takeaways**
- ✅ **CSV Reading/Writing** - Efficiently process CSV files with Go's encoding/csv
- ✅ **Struct Mapping** - Convert CSV data to strongly-typed Go structs
- ✅ **Streaming Processing** - Handle large files without memory issues
- ✅ **Data Analysis** - Build analytics and reporting systems
- ✅ **Error Handling** - Properly validate and handle malformed data

</div>

---

## 🚀 What's Next?

Ready to explore XML processing? Continue with:

<div style="text-align: center; margin: 30px 0;">
<a href="./module-8-3-xml-processing.md" style="background: linear-gradient(135deg, #4facfe 0%, #00f2fe 100%); color: white; padding: 15px 30px; border-radius: 25px; text-decoration: none; font-weight: bold; font-size: 18px;">🏷️ Next: XML Processing →</a>
</div>

---

*Happy coding! 🎉*
