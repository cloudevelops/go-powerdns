package powerdns

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	//"github.com/davecgh/go-spew/spew"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Powerdns struct {
	Hostname    string
	Apikey      string
	VerifySSL   bool
	BaseURL     string
	NameServers []string
	client      *http.Client
}

func NewPowerdns(HostName string, ApiKey string, NameServers []string) *Powerdns {
	var powerdns *Powerdns
	var tr *http.Transport

	powerdns = new(Powerdns)
	powerdns.Hostname = HostName
	powerdns.Apikey = ApiKey
	powerdns.VerifySSL = false
	powerdns.BaseURL = "http://" + powerdns.Hostname + ":8081/api/v1/servers/localhost/"
	powerdns.NameServers = NameServers
	if powerdns.VerifySSL {
		tr = &http.Transport{}
	} else {
		tr = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}
	powerdns.client = &http.Client{Transport: tr}
	return powerdns
}

func (powerdns *Powerdns) Post(endpoint string, jsonData []byte) (map[string]interface{}, error) {
	var target string
	var data interface{}
	var req *http.Request
	fmt.Println("PDNS.Post: endpoint: " + endpoint + " ,jsonData:" + string(jsonData))
	target = powerdns.BaseURL + endpoint
	req, err := http.NewRequest("POST", target, bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.Itoa(len(jsonData)))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-API-Key", powerdns.Apikey)
	r, err := powerdns.client.Do(req)
	defer r.Body.Close()
	if err != nil {
		fmt.Println("Error while posting")
		fmt.Println(err)
		return nil, err
	}
	if r.StatusCode < 200 || r.StatusCode > 299 {
		return nil, errors.New("HTTP Error " + r.Status)
	}

	response, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error while reading body")
		fmt.Println(err)
		return nil, err
	}
	err = json.Unmarshal(response, &data)
	if err != nil {
		fmt.Println("Error while processing JSON")
		fmt.Println(err)
		return nil, err
	}
	m := data.(map[string]interface{})
	return m, nil
}

func (powerdns *Powerdns) Get(endpoint string) (interface{}, error) {
	var target string
	var data interface{}
	fmt.Println("PDNS.Get: endpoint: " + endpoint)
	target = powerdns.BaseURL + endpoint
	req, err := http.NewRequest("GET", target, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-API-Key", powerdns.Apikey)
	r, err := powerdns.client.Do(req)
	defer r.Body.Close()
	if err != nil {
		fmt.Println("Error while getting")
		fmt.Println(err)
		return nil, err
	}
	if r.StatusCode < 200 || r.StatusCode > 299 {
		return nil, errors.New("HTTP Error " + r.Status)
	}

	response, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error while reading body")
		fmt.Println(err)
		return nil, err
	}
	err = json.Unmarshal(response, &data)
	if err != nil {
		fmt.Println("Error while processing JSON")
		fmt.Println(err)
		return nil, err
	}
	return data, nil
}

func (powerdns *Powerdns) Delete(endpoint string) (map[string]interface{}, error) {
	var target string
	//	var data interface{}
	var req *http.Request

	fmt.Println("PDNS.Delete: endpoint: " + endpoint)
	target = powerdns.BaseURL + endpoint
	req, err := http.NewRequest("DELETE", target, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-API-Key", powerdns.Apikey)
	r, err := powerdns.client.Do(req)
	defer r.Body.Close()
	if err != nil {
		fmt.Println("Error while deleting")
		fmt.Println(err)
		return nil, err
	}
	if r.StatusCode < 200 || r.StatusCode > 299 {
		return nil, errors.New("HTTP Error " + r.Status)
	}
	// Propably not needed here, as delete domain doesnt return anything just http response code
	//
	//	response, err := ioutil.ReadAll(r.Body)
	//if err != nil {
	//fmt.Println("Error while reading body")
	//fmt.Println(err)
	//return nil, err
	//}
	// Propably not needed here, as delete domain doesnt return anything just http response code
	//	err = json.Unmarshal(response, &data)
	//	if err != nil {
	//		fmt.Println("Error while processing JSON")
	//		fmt.Println(err)
	//		return nil, err
	//	}
	//m := data.(map[string]interface{})
	return nil, err
}

func (powerdns *Powerdns) Patch(endpoint string, jsonData []byte) (err error) {
	var target string
	var req *http.Request
	fmt.Println("PDNS.Patch: endpoint: " + endpoint + " ,jsonData:" + string(jsonData))
	target = powerdns.BaseURL + endpoint
	//fmt.Println("POST form " + target)
	req, err = http.NewRequest("PATCH", target, bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.Itoa(len(jsonData)))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-API-Key", powerdns.Apikey)
	r, err := powerdns.client.Do(req)
	defer r.Body.Close()
	if err != nil {
		fmt.Println("Error while patching")
		fmt.Println(err)
		return err
	}
	if r.StatusCode < 200 || r.StatusCode > 299 {
		return errors.New("HTTP Error " + r.Status)
	}
	return nil
}

func (powerdns *Powerdns) Put(endpoint string, jsonData []byte) (err error) {
	var target string
	var req *http.Request

	fmt.Println("PDNS.Put: endpoint: " + endpoint + " ,jsonData:" + string(jsonData))
	target = powerdns.BaseURL + endpoint
	//fmt.Println("POST form " + target)
	req, err = http.NewRequest("PUT", target, bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.Itoa(len(jsonData)))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-API-Key", powerdns.Apikey)
	r, err := powerdns.client.Do(req)
	defer r.Body.Close()
	if err != nil {
		fmt.Println("Error while patching")
		fmt.Println(err)
		return err
	}
	if r.StatusCode < 200 || r.StatusCode > 299 {
		return errors.New("HTTP Error " + r.Status)
	}
	return nil
}

type RrSet struct {
	Name       string        `json:"name"`
	DType      string        `json:"type"`
	Ttl        int           `json:"ttl"`
	ChangeType string        `json:"changetype"`
	Records    []interface{} `json:"records"`
}

type RrSlice []RrSet

type Record struct {
	Content  string `json:"content"`
	Disabled bool   `json:"disabled"`
	Name     string `json:"name"`
	Ttl      int    `json:"ttl"`
	DType    string `json:"type"`
}

type RecordSlice []Record

func (powerdns *Powerdns) UpdateRecord(domain string, dtype string, name string, content string, ttl int) error {

	var recordSlice []interface{}
	var rrSlice []interface{}
	fmt.Println("PDNS.UpdateRecord: Domain: " + domain + " ,dtype:" + dtype + " ,name:" + name + " ,content:" + content + " ,ttl:" + strconv.Itoa(ttl))
	record := Record{
		Content:  content,
		Disabled: false,
		Name:     name + "." + domain + ".",
		Ttl:      ttl,
		DType:    dtype,
	}
	recordSlice = append(recordSlice, record)
	rrSet := RrSet{
		Name:       name + "." + domain + ".",
		DType:      dtype,
		Ttl:        ttl,
		ChangeType: "REPLACE",
		Records:    recordSlice,
	}
	rrSlice = append(rrSlice, rrSet)
	update := make(map[string]interface{})
	update["rrsets"] = rrSlice
	jsonText, err := json.Marshal(update)
	topDomain, err := powerdns.GetTopDomain(domain)
	if err != nil {
		fmt.Println("PDNS.UpdateRecord: Could not get topdomain for " + domain + ", reverting to domain")
		fmt.Println(err)
		topDomain = domain
	} else {
		fmt.Println("PDNS.UpdateRecord: Got topdomain " + topDomain + " for domain " + domain)
	}
	_, err = powerdns.Get("zones/" + domain)
	if err != nil {
		fmt.Println("PDNS.UpdateRecord: Domain " + domain + " not found, attempting to create it")
		err := powerdns.CreateDomain(domain)
		if err != nil {
			fmt.Println("PDNS.UpdateRecord: Failed to create domain:" + domain)
			return err
		}
		// create a record in just created domain
		err = powerdns.Patch("zones/"+domain, jsonText)
		if err != nil {
			fmt.Println("PDNS.UpdateRecord: Error updating record at path zones/" + domain + " ,content:" + string(jsonText))
			fmt.Println(err)
			return err
		}
		return nil
	} else {
		// create an record in topdomain
		err = powerdns.Patch("zones/"+topDomain, jsonText)
		if err != nil {
			fmt.Println("PDNS.UpdateRecord: Error updating record at path zones/" + topDomain + " ,content:" + string(jsonText))
			fmt.Println(err)
			return err
		}
		return nil
	}
}

func (powerdns *Powerdns) UpdateRec(domain string, dtype string, name string, content string, ttl int) error {

	var recordSlice []interface{}
	var rrSlice []interface{}
	fmt.Println("PDNS.UpdateRec: Domain: " + domain + " ,dtype:" + dtype + " ,name:" + name + " ,content:" + content + " ,ttl:" + strconv.Itoa(ttl))
	record := Record{
		Content:  content,
		Disabled: false,
		Name:     name,
		Ttl:      ttl,
		DType:    dtype,
	}
	recordSlice = append(recordSlice, record)
	rrSet := RrSet{
		Name:       name,
		DType:      dtype,
		Ttl:        ttl,
		ChangeType: "REPLACE",
		Records:    recordSlice,
	}
	rrSlice = append(rrSlice, rrSet)
	update := make(map[string]interface{})
	update["rrsets"] = rrSlice
	jsonText, err := json.Marshal(update)
	topDomain, err := powerdns.GetTopDomain(domain)
	if err != nil {
		fmt.Println("PDNS.UpdateRec: Could not get topdomain, reverting to domain: " + domain)
		fmt.Println(err)
		topDomain = domain
	}
	err = powerdns.Patch("zones/"+topDomain, jsonText)
	if err != nil {
		fmt.Println("PDNS.UpdateRec: Error updating record at path zones/" + topDomain + " ,content:" + string(jsonText))
		fmt.Println(err)
		return err
	}
	return nil
}

func (powerdns *Powerdns) GetTopDomain(domain string) (topdomain string, err error) {
	fmt.Println("PDNS.GetTopDomain: Domain: " + domain)
	topSlice := strings.Split(domain, ".")
	for i := 0; i < len(topSlice); i++ {
		topdomain = ""
		for n := i; n < len(topSlice); n++ {
			topdomain = topdomain + topSlice[n] + "."
		}
		topDomainData, err := powerdns.Get("zones/" + topdomain)
		if err == nil {
			topDomainDataMap := topDomainData.(map[string]interface{})
			if topDomainDataMap["soa_edit_api"] != "INCEPTION-INCREMENT" {
				fmt.Println("PDNS.GetTopDomain: Updating soa_edit_api and soa_edit at path zones/" + topdomain)
				update := make(map[string]string)
				update["soa_edit_api"] = "INCEPTION-INCREMENT"
				update["soa_edit"] = "INCEPTION-INCREMENT"
				update["kind"] = "Master"
				jsonText, err := json.Marshal(update)
				err = powerdns.Put("zones/"+topdomain, jsonText)
				if err != nil {
					fmt.Println("PDNS.GetTopDomain: Error updating soa_edit_api and soa_edit at path zones/" + topdomain + " ,content:" + string(jsonText))
					fmt.Println(err)
				}
			}
			return topdomain, err
		}
	}
	return topdomain, errors.New("PDNS.GetTopDomain: Did not found domain:" + domain + " for topdomain:" + topdomain)
}

func (powerdns *Powerdns) DeleteRecord(domain string, dtype string, name string) error {

	var recordSlice []interface{}
	var rrSlice []interface{}
	fmt.Println("PDNS.DeleteRecord: Domain: " + domain + " ,dtype:" + dtype + " ,name:" + name)
	record := Record{
		Disabled: false,
		Name:     name + "." + domain + ".",
		DType:    dtype,
	}
	recordSlice = append(recordSlice, record)
	rrSet := RrSet{
		Name:       name + "." + domain + ".",
		DType:      dtype,
		ChangeType: "DELETE",
		Records:    recordSlice,
	}
	rrSlice = append(rrSlice, rrSet)
	update := make(map[string]interface{})
	update["rrsets"] = rrSlice
	jsonText, err := json.Marshal(update)
	topDomain, err := powerdns.GetTopDomain(domain)
	if err != nil {
		fmt.Println("PDNS.DeleteRecord: Could not get topdomain, reverting to domain: " + domain)
		fmt.Println(err)
		topDomain = domain
	}
	err = powerdns.Patch("zones/"+topDomain, jsonText)
	if err != nil {
		fmt.Println("PDNS.DeleteRecord: Error updating record at path zones/" + topDomain + " ,content:" + string(jsonText))
		fmt.Println(err)
		return err
	}
	return nil
}

func (powerdns *Powerdns) DeleteRec(domain string, dtype string, name string) error {

	var recordSlice []interface{}
	var rrSlice []interface{}
	fmt.Println("PDNS.DeleteRec: Domain: " + domain + " ,dtype:" + dtype + " ,name:" + name)
	record := Record{
		Disabled: false,
		Name:     name,
		DType:    dtype,
	}
	recordSlice = append(recordSlice, record)
	rrSet := RrSet{
		Name:       name,
		DType:      dtype,
		ChangeType: "DELETE",
		Records:    recordSlice,
	}
	rrSlice = append(rrSlice, rrSet)
	update := make(map[string]interface{})
	update["rrsets"] = rrSlice
	jsonText, err := json.Marshal(update)
	topDomain, err := powerdns.GetTopDomain(domain)
	if err != nil {
		fmt.Println("PDNS.DeleteRec: Could not get topdomain, reverting to domain: " + domain)
		fmt.Println(err)
		topDomain = domain
	}
	err = powerdns.Patch("zones/"+topDomain, jsonText)
	if err != nil {
		fmt.Println("PDNS.DeleteRec: Error updating record at path zones/" + topDomain + " ,content:" + string(jsonText))
		fmt.Println(err)
		return err
	}
	return nil
}

func (powerdns *Powerdns) CreateDomain(domain string) error {
	// create domain itself
	type Domain struct {
		Name        string   `json:"name"`
		Kind        string   `json:"kind"`
		Masters     []string `json:"masters"`
		Nameservers []string `json:"nameservers"`
	}
	fmt.Println("PDNS.CreateDomain: domain: " + domain)
	if (domain == "") || (domain == ".") {
		fmt.Println("PDNS.CreateDomain: Domain in invalid format, refusing to create: " + domain)
		return errors.New("Domain in invalid format")
	}
	masters := make([]string, 0)
	var CanonicalNameServers []string
	for _, nameserver := range powerdns.NameServers {
		canonicalNameServer := nameserver + "."
		CanonicalNameServers = append(CanonicalNameServers, canonicalNameServer)
	}
	canonicalDomain := domain + "."
	domainSet := Domain{
		Name:        canonicalDomain,
		Kind:        "Master",
		Masters:     masters,
		Nameservers: CanonicalNameServers,
	}

	jsonText, err := json.Marshal(domainSet)

	_, err = powerdns.Post("zones", jsonText)
	if err != nil {
		fmt.Println("PDNS.CreateDomain: Error creating domain: " + domain)
		fmt.Println(err)
		return err
	}
	// initialize SOA record
	t := time.Now()
	timestamp := t.Format("20060102") + "01"
	soa := CanonicalNameServers[0] + " hostmaster. " + timestamp + " 28800 7200 604800 86400"
	err = powerdns.UpdateRec(canonicalDomain, "SOA", canonicalDomain, soa, 10)
	if err != nil {
		fmt.Println("PDNS.CreateDomain: Failed to update SOA record, domain: " + canonicalDomain + ", name: " + canonicalDomain + ", content: " + soa + " !")
	}
	fmt.Println("PDNS.CreateDomain: Updated SOA record, domain: " + canonicalDomain + ", name: " + canonicalDomain + ", content: " + soa + " !")

	return nil
}

func (powerdns *Powerdns) DeleteDomain(domain string) error {

	fmt.Println("PDNS.DeleteDomain: Domain: " + domain)
	_, err := powerdns.Delete("zones/" + domain)
	if err != nil {
		fmt.Println("PDNS.DeleteDomain: Error deleting domain at path zones/" + domain)
		fmt.Println(err)
		return err
	}
	return nil
}
