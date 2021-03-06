package asciidocgo

// Asciidoc Document, onced loaded from an IO, string or array
type Document struct {
	monitorData *monitorData
	data        []string
	options     map[string]string
}

type monitorData struct {
	readTime   int
	parseTime  int
	renderTime int
	writeTime  int
}

// Error returned when accessing times on a Document not monitored
type NotMonitoredError struct {
	msg string // description of error
}

// Print description of a non-monitored error
func (e *NotMonitoredError) Error() string { return e.msg }

// Check if a Document is supposed to be monitored
func (d *Document) IsMonitored() bool {
	return (d.monitorData != nil)
}

// Setup a monitor for the document
// (or does nothing if the monitor already exists).
// Returns self, for easy composition
func (d *Document) Monitor() *Document {
	if d.monitorData == nil {
		d.monitorData = new(monitorData)
	}
	return d
}

/*
Initialize an Asciidoc object.
- data: The Array of Strings holding the Asciidoc source document. (default: [])
- options - A Hash of options to control processing, such as setting the safe mode (:safe), suppressing the header/footer (:header_footer) and attribute overrides (:attributes)
(default: {})

Examples

    data = File.readlines(filename)
    doc  = Asciidoctor::Document.new(data)
    puts doc.render
*/
func NewDocument(data []string, options map[string]string) *Document {
	document := &Document{nil, data, options}
	return document
}

// Time to read the document from IO source
// Error if document didn't activated the monitoring
func (d *Document) ReadTime() (readTime int, err error) {
	if d.IsMonitored() == false {
		return 0, &NotMonitoredError{"No readTime: current document is not monitored"}
	}
	return d.monitorData.readTime, nil
}

// Time to parse the document once read from IO source
// Error if document didn't activated the monitoring
func (d *Document) ParseTime() (parseTime int, err error) {
	if d.IsMonitored() == false {
		return 0, &NotMonitoredError{"No parseTime: current document is not monitored"}
	}
	return d.monitorData.parseTime, nil
}

// Load means Read plus Parse times
// Error if document didn't activated the monitoring
func (d *Document) LoadTime() (loadTime int, err error) {
	if d.IsMonitored() == false {
		return 0, &NotMonitoredError{"No loadTime: current document is not monitored"}
	}
	readTime, _ := d.ReadTime()
	parseTime, _ := d.ParseTime()
	return readTime + parseTime, nil
}

// Time to render the document once loaded
// Error if document didn't activated the monitoring
func (d *Document) RenderTime() (renderTime int, err error) {
	if d.IsMonitored() == false {
		return 0, &NotMonitoredError{"No ploadTime: current document is not monitored"}
	}
	return d.monitorData.renderTime, nil
}

// LoadRender means Load plus Render times
// Error if document didn't activated the monitoring
func (d *Document) LoadRenderTime() (loadRenderTime int, err error) {
	if d.IsMonitored() == false {
		return 0, &NotMonitoredError{"No loadTime: current document is not monitored"}
	}
	loadTime, _ := d.LoadTime()
	renderTime, _ := d.RenderTime()
	return loadTime + renderTime, nil
}

// Time to write the document once rendered
// Error if document didn't activated the monitoring
func (d *Document) WriteTime() (writeTime int, err error) {
	if d.IsMonitored() == false {
		return 0, &NotMonitoredError{"No ploadTime: current document is not monitored"}
	}
	return d.monitorData.writeTime, nil
}

// Total means LoadRender plus Write times
// Error if document didn't activated the monitoring
func (d *Document) TotalTime() (totalTime int, err error) {
	if d.IsMonitored() == false {
		return 0, &NotMonitoredError{"No loadTime: current document is not monitored"}
	}
	loadRenderTime, _ := d.LoadRenderTime()
	writeTime, _ := d.WriteTime()
	return loadRenderTime + writeTime, nil
}
