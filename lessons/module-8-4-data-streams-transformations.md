# 🚀 Module 8.4: Data Streams and Transformations

<div align="center">

![Go Logo](https://golang.org/lib/godoc/images/go-logo-blue.svg)

**Master Data Streaming and Transformation Pipelines**

*From Zero to Go Hero - Module 8.4*

</div>

---

## 📋 Lesson Overview

| **Duration** | **Difficulty** | **Prerequisites** |
|--------------|----------------|-------------------|
| 1.5-2 hours  | Advanced       | Concurrency, Channels, File I/O, JSON/CSV/XML |

### 🎯 Learning Objectives
By the end of this lesson, you will:
- ✅ Build streaming data processing pipelines
- ✅ Transform data between different formats efficiently
- ✅ Handle large datasets with minimal memory usage
- ✅ Implement concurrent data processing patterns
- ✅ Create real-world ETL (Extract, Transform, Load) systems

---

## 🌟 What are Data Streams?

<div style="background: linear-gradient(135deg, #fa709a 0%, #fee140 100%); padding: 20px; border-radius: 10px; color: white; margin: 20px 0;">

### 🌊 **Data Streaming Concepts**

Data streaming is the process of handling data as a continuous flow rather than loading everything into memory at once. This approach is essential for processing large datasets, real-time data, and building scalable applications.

**Think of data streams like a river** - instead of trying to hold all the water at once, you process it as it flows by!

</div>

### 🏗️ **Stream Processing Pipeline**
```
┌─────────────────────────────────────────────────────────┐
│                Data Stream Pipeline                     │
├─────────────────────────────────────────────────────────┤
│  Input → Transform → Filter → Aggregate → Output       │
│    ↓         ↓         ↓         ↓         ↓           │
│  [Data]   [Clean]   [Valid]   [Group]   [Result]       │
│  Source   Format    Records   Process   Storage        │
└─────────────────────────────────────────────────────────┘
```

---

## 🔧 Building Stream Processing Pipelines

<div style="background: #f8f9fa; border-left: 4px solid #28a745; padding: 15px; margin: 10px 0;">

**Core Concepts** - Stream processing involves:

```go
// Data flows through channels between processing stages
Input Channel → Transform → Filter → Output Channel

// Each stage can run concurrently
go processStage1(inputCh, outputCh1)
go processStage2(outputCh1, outputCh2)
go processStage3(outputCh2, finalCh)
```

</div>

### 🎯 **Basic Stream Pipeline**

```go
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "math/rand"
    "strconv"
    "strings"
    "sync"
    "time"
)

// DataRecord represents a generic data record
type DataRecord struct {
    ID        int                    `json:"id"`
    Timestamp time.Time              `json:"timestamp"`
    Type      string                 `json:"type"`
    Data      map[string]interface{} `json:"data"`
}

// StreamProcessor handles data stream processing
type StreamProcessor struct {
    inputCh    chan DataRecord
    outputCh   chan DataRecord
    errorCh    chan error
    wg         sync.WaitGroup
    processors int
}

// NewStreamProcessor creates a new stream processor
func NewStreamProcessor(bufferSize, processors int) *StreamProcessor {
    return &StreamProcessor{
        inputCh:    make(chan DataRecord, bufferSize),
        outputCh:   make(chan DataRecord, bufferSize),
        errorCh:    make(chan error, bufferSize),
        processors: processors,
    }
}

// Start begins the stream processing pipeline
func (sp *StreamProcessor) Start() {
    // Start multiple processor goroutines
    for i := 0; i < sp.processors; i++ {
        sp.wg.Add(1)
        go sp.processWorker(i)
    }
    
    fmt.Printf("🚀 Started %d stream processors\n", sp.processors)
}

// processWorker processes data records
func (sp *StreamProcessor) processWorker(workerID int) {
    defer sp.wg.Done()
    
    for record := range sp.inputCh {
        // Transform the record
        transformed, err := sp.transformRecord(record, workerID)
        if err != nil {
            sp.errorCh <- fmt.Errorf("worker %d: %w", workerID, err)
            continue
        }
        
        // Validate the record
        if sp.validateRecord(transformed) {
            sp.outputCh <- transformed
        }
    }
}

// transformRecord applies transformations to a data record
func (sp *StreamProcessor) transformRecord(record DataRecord, workerID int) (DataRecord, error) {
    // Add processing metadata
    if record.Data == nil {
        record.Data = make(map[string]interface{})
    }
    
    record.Data["processed_by"] = fmt.Sprintf("worker-%d", workerID)
    record.Data["processed_at"] = time.Now()
    
    // Normalize the type field
    record.Type = strings.ToUpper(strings.TrimSpace(record.Type))
    
    // Add computed fields based on type
    switch record.Type {
    case "SENSOR":
        if temp, exists := record.Data["temperature"]; exists {
            if tempFloat, ok := temp.(float64); ok {
                record.Data["temp_fahrenheit"] = tempFloat*9/5 + 32
                record.Data["temp_category"] = sp.categorizeTemperature(tempFloat)
            }
        }
    case "USER":
        if age, exists := record.Data["age"]; exists {
            if ageFloat, ok := age.(float64); ok {
                record.Data["age_group"] = sp.categorizeAge(int(ageFloat))
            }
        }
    case "TRANSACTION":
        if amount, exists := record.Data["amount"]; exists {
            if amountFloat, ok := amount.(float64); ok {
                record.Data["amount_category"] = sp.categorizeAmount(amountFloat)
            }
        }
    }
    
    return record, nil
}

// validateRecord checks if a record is valid
func (sp *StreamProcessor) validateRecord(record DataRecord) bool {
    // Basic validation rules
    if record.ID <= 0 {
        return false
    }
    
    if record.Type == "" {
        return false
    }
    
    if record.Timestamp.IsZero() {
        return false
    }
    
    return true
}

// Helper functions for categorization
func (sp *StreamProcessor) categorizeTemperature(temp float64) string {
    switch {
    case temp < 0:
        return "freezing"
    case temp < 10:
        return "cold"
    case temp < 20:
        return "cool"
    case temp < 30:
        return "warm"
    default:
        return "hot"
    }
}

func (sp *StreamProcessor) categorizeAge(age int) string {
    switch {
    case age < 18:
        return "minor"
    case age < 30:
        return "young_adult"
    case age < 50:
        return "adult"
    case age < 65:
        return "middle_aged"
    default:
        return "senior"
    }
}

func (sp *StreamProcessor) categorizeAmount(amount float64) string {
    switch {
    case amount < 10:
        return "small"
    case amount < 100:
        return "medium"
    case amount < 1000:
        return "large"
    default:
        return "very_large"
    }
}

// SendRecord sends a record to the input channel
func (sp *StreamProcessor) SendRecord(record DataRecord) {
    sp.inputCh <- record
}

// GetOutputChannel returns the output channel
func (sp *StreamProcessor) GetOutputChannel() <-chan DataRecord {
    return sp.outputCh
}

// GetErrorChannel returns the error channel
func (sp *StreamProcessor) GetErrorChannel() <-chan error {
    return sp.errorCh
}

// Stop stops the stream processor
func (sp *StreamProcessor) Stop() {
    close(sp.inputCh)
    sp.wg.Wait()
    close(sp.outputCh)
    close(sp.errorCh)
    fmt.Println("🛑 Stream processor stopped")
}

// DataGenerator generates sample data for testing
type DataGenerator struct {
    recordID int
}

// NewDataGenerator creates a new data generator
func NewDataGenerator() *DataGenerator {
    return &DataGenerator{recordID: 1}
}

// GenerateRecord creates a sample data record
func (dg *DataGenerator) GenerateRecord() DataRecord {
    recordTypes := []string{"SENSOR", "USER", "TRANSACTION"}
    recordType := recordTypes[rand.Intn(len(recordTypes))]
    
    record := DataRecord{
        ID:        dg.recordID,
        Timestamp: time.Now(),
        Type:      recordType,
        Data:      make(map[string]interface{}),
    }
    
    // Generate type-specific data
    switch recordType {
    case "SENSOR":
        record.Data["temperature"] = rand.Float64()*50 - 10 // -10 to 40 degrees
        record.Data["humidity"] = rand.Float64() * 100      // 0 to 100%
        record.Data["sensor_id"] = fmt.Sprintf("TEMP_%03d", rand.Intn(100))
    case "USER":
        record.Data["age"] = float64(rand.Intn(80) + 18) // 18 to 98 years
        record.Data["country"] = []string{"USA", "UK", "Canada", "Germany", "Japan"}[rand.Intn(5)]
        record.Data["user_id"] = fmt.Sprintf("USER_%06d", rand.Intn(1000000))
    case "TRANSACTION":
        record.Data["amount"] = rand.Float64() * 2000 // 0 to 2000
        record.Data["currency"] = []string{"USD", "EUR", "GBP", "JPY"}[rand.Intn(4)]
        record.Data["merchant"] = fmt.Sprintf("MERCHANT_%03d", rand.Intn(500))
    }
    
    dg.recordID++
    return record
}

func main() {
    // Create stream processor
    processor := NewStreamProcessor(100, 3) // Buffer size 100, 3 workers
    processor.Start()
    
    // Create data generator
    generator := NewDataGenerator()
    
    // Statistics tracking
    var processedCount, errorCount int
    
    // Start result collector goroutine
    go func() {
        for record := range processor.GetOutputChannel() {
            processedCount++
            
            // Display every 10th record
            if processedCount%10 == 0 {
                fmt.Printf("📊 Processed %d records\n", processedCount)
                
                // Show sample record
                if processedCount%50 == 0 {
                    jsonData, _ := json.MarshalIndent(record, "", "  ")
                    fmt.Printf("📄 Sample Record:\n%s\n\n", string(jsonData))
                }
            }
        }
    }()
    
    // Start error collector goroutine
    go func() {
        for err := range processor.GetErrorChannel() {
            errorCount++
            fmt.Printf("❌ Error %d: %v\n", errorCount, err)
        }
    }()
    
    // Generate and send data
    fmt.Println("🔄 Generating and processing data...")
    start := time.Now()
    
    for i := 0; i < 200; i++ {
        record := generator.GenerateRecord()
        processor.SendRecord(record)
        
        // Small delay to simulate real-time data
        time.Sleep(10 * time.Millisecond)
    }
    
    // Stop processing
    processor.Stop()
    
    // Wait a bit for final processing
    time.Sleep(100 * time.Millisecond)
    
    duration := time.Since(start)
    fmt.Printf("\n✅ Processing completed!\n")
    fmt.Printf("📈 Statistics:\n")
    fmt.Printf("  - Total processed: %d records\n", processedCount)
    fmt.Printf("  - Errors: %d\n", errorCount)
    fmt.Printf("  - Duration: %v\n", duration)
    fmt.Printf("  - Rate: %.2f records/second\n", float64(processedCount)/duration.Seconds())
}
```

---

## 🔄 Format Transformation Pipeline

<div style="background: linear-gradient(135deg, #667eea 0%, #764ba2 100%); padding: 20px; border-radius: 10px; color: white; margin: 20px 0;">

### 🎨 **Multi-Format Data Converter**

Let's build a system that can convert data between JSON, CSV, and XML formats while applying transformations.

</div>

### 🔧 **Format Converter Implementation**

```go
package main

import (
    "encoding/csv"
    "encoding/json"
    "encoding/xml"
    "fmt"
    "io"
    "os"
    "reflect"
    "strconv"
    "strings"
    "sync"
)

// DataFormat represents supported data formats
type DataFormat int

const (
    FormatJSON DataFormat = iota
    FormatCSV
    FormatXML
)

// ConversionJob represents a data conversion task
type ConversionJob struct {
    ID           int
    InputFormat  DataFormat
    OutputFormat DataFormat
    InputData    []byte
    OutputData   []byte
    Error        error
}

// DataConverter handles format conversions
type DataConverter struct {
    workers int
    jobCh   chan ConversionJob
    resultCh chan ConversionJob
    wg      sync.WaitGroup
}

// NewDataConverter creates a new data converter
func NewDataConverter(workers int) *DataConverter {
    return &DataConverter{
        workers:  workers,
        jobCh:    make(chan ConversionJob, 100),
        resultCh: make(chan ConversionJob, 100),
    }
}

// Start begins the conversion workers
func (dc *DataConverter) Start() {
    for i := 0; i < dc.workers; i++ {
        dc.wg.Add(1)
        go dc.worker(i)
    }
    fmt.Printf("🔄 Started %d conversion workers\n", dc.workers)
}

// worker processes conversion jobs
func (dc *DataConverter) worker(workerID int) {
    defer dc.wg.Done()
    
    for job := range dc.jobCh {
        fmt.Printf("👷 Worker %d processing job %d: %s → %s\n", 
            workerID, job.ID, 
            dc.formatName(job.InputFormat), 
            dc.formatName(job.OutputFormat))
        
        // Parse input data
        data, err := dc.parseInput(job.InputData, job.InputFormat)
        if err != nil {
            job.Error = fmt.Errorf("parse error: %w", err)
            dc.resultCh <- job
            continue
        }
        
        // Apply transformations
        transformedData := dc.transformData(data)
        
        // Generate output
        output, err := dc.generateOutput(transformedData, job.OutputFormat)
        if err != nil {
            job.Error = fmt.Errorf("output error: %w", err)
            dc.resultCh <- job
            continue
        }
        
        job.OutputData = output
        dc.resultCh <- job
    }
}

// parseInput parses input data based on format
func (dc *DataConverter) parseInput(data []byte, format DataFormat) ([]map[string]interface{}, error) {
    switch format {
    case FormatJSON:
        return dc.parseJSON(data)
    case FormatCSV:
        return dc.parseCSV(data)
    case FormatXML:
        return dc.parseXML(data)
    default:
        return nil, fmt.Errorf("unsupported input format")
    }
}

// parseJSON parses JSON data
func (dc *DataConverter) parseJSON(data []byte) ([]map[string]interface{}, error) {
    var result []map[string]interface{}
    
    // Try parsing as array first
    if err := json.Unmarshal(data, &result); err == nil {
        return result, nil
    }
    
    // Try parsing as single object
    var single map[string]interface{}
    if err := json.Unmarshal(data, &single); err == nil {
        return []map[string]interface{}{single}, nil
    }
    
    return nil, fmt.Errorf("invalid JSON format")
}

// parseCSV parses CSV data
func (dc *DataConverter) parseCSV(data []byte) ([]map[string]interface{}, error) {
    reader := csv.NewReader(strings.NewReader(string(data)))
    
    // Read header
    header, err := reader.Read()
    if err != nil {
        return nil, fmt.Errorf("failed to read CSV header: %w", err)
    }
    
    var result []map[string]interface{}
    
    // Read records
    for {
        record, err := reader.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            return nil, fmt.Errorf("failed to read CSV record: %w", err)
        }
        
        // Convert record to map
        recordMap := make(map[string]interface{})
        for i, value := range record {
            if i < len(header) {
                // Try to convert to appropriate type
                recordMap[header[i]] = dc.convertValue(value)
            }
        }
        result = append(result, recordMap)
    }
    
    return result, nil
}

// parseXML parses XML data (simplified)
func (dc *DataConverter) parseXML(data []byte) ([]map[string]interface{}, error) {
    // For this example, we'll assume a simple XML structure
    // In practice, you'd need more sophisticated XML parsing
    
    type XMLRecord struct {
        XMLName xml.Name               `xml:"record"`
        Fields  map[string]interface{} `xml:",any"`
    }
    
    type XMLRoot struct {
        XMLName xml.Name    `xml:"root"`
        Records []XMLRecord `xml:"record"`
    }
    
    var root XMLRoot
    if err := xml.Unmarshal(data, &root); err != nil {
        return nil, fmt.Errorf("failed to parse XML: %w", err)
    }
    
    var result []map[string]interface{}
    for _, record := range root.Records {
        result = append(result, record.Fields)
    }
    
    return result, nil
}

// convertValue attempts to convert string to appropriate type
func (dc *DataConverter) convertValue(value string) interface{} {
    value = strings.TrimSpace(value)
    
    // Try boolean
    if strings.ToLower(value) == "true" {
        return true
    }
    if strings.ToLower(value) == "false" {
        return false
    }
    
    // Try integer
    if intVal, err := strconv.Atoi(value); err == nil {
        return intVal
    }
    
    // Try float
    if floatVal, err := strconv.ParseFloat(value, 64); err == nil {
        return floatVal
    }
    
    // Return as string
    return value
}

// transformData applies transformations to the data
func (dc *DataConverter) transformData(data []map[string]interface{}) []map[string]interface{} {
    var result []map[string]interface{}
    
    for _, record := range data {
        transformed := make(map[string]interface{})
        
        // Copy and transform fields
        for key, value := range record {
            // Normalize field names (lowercase, replace spaces with underscores)
            normalizedKey := strings.ToLower(strings.ReplaceAll(key, " ", "_"))
            
            // Apply value transformations
            transformed[normalizedKey] = dc.transformValue(normalizedKey, value)
        }
        
        // Add metadata
        transformed["_transformed_at"] = "2024-01-15T10:30:00Z"
        transformed["_record_id"] = len(result) + 1
        
        result = append(result, transformed)
    }
    
    return result
}

// transformValue applies transformations to individual values
func (dc *DataConverter) transformValue(key string, value interface{}) interface{} {
    // Apply key-specific transformations
    switch key {
    case "email":
        if str, ok := value.(string); ok {
            return strings.ToLower(str)
        }
    case "name", "first_name", "last_name":
        if str, ok := value.(string); ok {
            return strings.Title(strings.ToLower(str))
        }
    case "phone":
        if str, ok := value.(string); ok {
            // Remove non-numeric characters
            phone := strings.ReplaceAll(str, "-", "")
            phone = strings.ReplaceAll(phone, " ", "")
            phone = strings.ReplaceAll(phone, "(", "")
            phone = strings.ReplaceAll(phone, ")", "")
            return phone
        }
    }
    
    return value
}

// generateOutput generates output in the specified format
func (dc *DataConverter) generateOutput(data []map[string]interface{}, format DataFormat) ([]byte, error) {
    switch format {
    case FormatJSON:
        return dc.generateJSON(data)
    case FormatCSV:
        return dc.generateCSV(data)
    case FormatXML:
        return dc.generateXML(data)
    default:
        return nil, fmt.Errorf("unsupported output format")
    }
}

// generateJSON generates JSON output
func (dc *DataConverter) generateJSON(data []map[string]interface{}) ([]byte, error) {
    return json.MarshalIndent(data, "", "  ")
}

// generateCSV generates CSV output
func (dc *DataConverter) generateCSV(data []map[string]interface{}) ([]byte, error) {
    if len(data) == 0 {
        return []byte{}, nil
    }
    
    var output strings.Builder
    writer := csv.NewWriter(&output)
    
    // Get all unique field names
    fieldSet := make(map[string]bool)
    for _, record := range data {
        for field := range record {
            fieldSet[field] = true
        }
    }
    
    // Create sorted field list
    var fields []string
    for field := range fieldSet {
        fields = append(fields, field)
    }
    
    // Write header
    writer.Write(fields)
    
    // Write records
    for _, record := range data {
        var row []string
        for _, field := range fields {
            if value, exists := record[field]; exists {
                row = append(row, fmt.Sprintf("%v", value))
            } else {
                row = append(row, "")
            }
        }
        writer.Write(row)
    }
    
    writer.Flush()
    return []byte(output.String()), nil
}

// generateXML generates XML output
func (dc *DataConverter) generateXML(data []map[string]interface{}) ([]byte, error) {
    var output strings.Builder
    output.WriteString(`<?xml version="1.0" encoding="UTF-8"?>` + "\n")
    output.WriteString("<root>\n")
    
    for i, record := range data {
        output.WriteString(fmt.Sprintf("  <record id=\"%d\">\n", i+1))
        
        for key, value := range record {
            // Sanitize XML element name
            xmlKey := strings.ReplaceAll(key, " ", "_")
            xmlKey = strings.ReplaceAll(xmlKey, "-", "_")
            
            output.WriteString(fmt.Sprintf("    <%s>%v</%s>\n", xmlKey, value, xmlKey))
        }
        
        output.WriteString("  </record>\n")
    }
    
    output.WriteString("</root>\n")
    return []byte(output.String()), nil
}

// formatName returns the name of a format
func (dc *DataConverter) formatName(format DataFormat) string {
    switch format {
    case FormatJSON:
        return "JSON"
    case FormatCSV:
        return "CSV"
    case FormatXML:
        return "XML"
    default:
        return "Unknown"
    }
}

// SubmitJob submits a conversion job
func (dc *DataConverter) SubmitJob(job ConversionJob) {
    dc.jobCh <- job
}

// GetResults returns the results channel
func (dc *DataConverter) GetResults() <-chan ConversionJob {
    return dc.resultCh
}

// Stop stops the converter
func (dc *DataConverter) Stop() {
    close(dc.jobCh)
    dc.wg.Wait()
    close(dc.resultCh)
    fmt.Println("🛑 Data converter stopped")
}

func main() {
    // Create converter
    converter := NewDataConverter(2)
    converter.Start()
    
    // Sample data in different formats
    jsonData := `[
        {"name": "John Doe", "age": 30, "email": "JOHN@EXAMPLE.COM", "phone": "(555) 123-4567"},
        {"name": "jane smith", "age": 25, "email": "jane@example.com", "phone": "555-987-6543"}
    ]`
    
    csvData := `Name,Age,Email,Phone
John Doe,30,JOHN@EXAMPLE.COM,(555) 123-4567
jane smith,25,jane@example.com,555-987-6543`
    
    // Submit conversion jobs
    jobs := []ConversionJob{
        {ID: 1, InputFormat: FormatJSON, OutputFormat: FormatCSV, InputData: []byte(jsonData)},
        {ID: 2, InputFormat: FormatCSV, OutputFormat: FormatJSON, InputData: []byte(csvData)},
        {ID: 3, InputFormat: FormatJSON, OutputFormat: FormatXML, InputData: []byte(jsonData)},
    }
    
    // Submit jobs
    for _, job := range jobs {
        converter.SubmitJob(job)
    }
    
    // Collect results
    var completed int
    for completed < len(jobs) {
        result := <-converter.GetResults()
        completed++
        
        fmt.Printf("\n✅ Job %d completed\n", result.ID)
        if result.Error != nil {
            fmt.Printf("❌ Error: %v\n", result.Error)
        } else {
            fmt.Printf("📄 Output (%s):\n%s\n", 
                converter.formatName(result.OutputFormat), 
                string(result.OutputData))
        }
    }
    
    converter.Stop()
    fmt.Println("\n🎉 All conversions completed!")
}
```

---

## 🎯 Practice Exercises

### 🏋️ **Exercise 1: Real-time Log Processor**
Create a system that:
1. Processes log files in real-time
2. Filters by severity levels
3. Aggregates statistics by time windows
4. Exports alerts in multiple formats

### 🏋️ **Exercise 2: Data ETL Pipeline**
Build an ETL system that:
1. Extracts data from multiple sources (APIs, files, databases)
2. Transforms and validates the data
3. Loads results into different destinations
4. Handles errors and retries gracefully

### 🏋️ **Exercise 3: Stream Analytics Engine**
Create an analytics engine that:
1. Processes streaming data from multiple sources
2. Calculates real-time metrics and KPIs
3. Detects anomalies and patterns
4. Generates reports and dashboards

---

## 🎉 Summary

Congratulations! You've mastered data streams and transformations in Go. Here's what you've learned:

<div style="background: linear-gradient(135deg, #fa709a 0%, #fee140 100%); padding: 20px; border-radius: 10px; color: white; margin: 20px 0;">

### 🎯 **Key Takeaways**
- ✅ **Stream Processing** - Handle continuous data flows efficiently
- ✅ **Concurrent Pipelines** - Build scalable processing systems
- ✅ **Format Conversion** - Transform data between JSON, CSV, and XML
- ✅ **Memory Efficiency** - Process large datasets without memory issues
- ✅ **Real-world Applications** - Build ETL systems and data processors

</div>

---

## 🚀 Module 8 Complete!

You've successfully completed Module 8: Data Processing and Serialization! You now have the skills to:

- **Process JSON, CSV, and XML** data efficiently
- **Build streaming data pipelines** for large-scale processing
- **Transform data** between different formats
- **Handle real-world data processing** challenges

<div style="text-align: center; margin: 30px 0;">
<a href="../module-9-tooling-dependencies-testing.md" style="background: linear-gradient(135deg, #667eea 0%, #764ba2 100%); color: white; padding: 15px 30px; border-radius: 25px; text-decoration: none; font-weight: bold; font-size: 18px;">🛠️ Continue to Module 9: Tooling & Testing →</a>
</div>

---

*Happy coding! 🎉*
