package tests

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/priyanshu360/lab-rank/dashboard/models"
)

type endpointTest struct {
	name         string
	path         string
	method       string
	body         string
	header       map[string]string
	wantRespCode int
	// wantRespBody   string
	// wantRespHeader []byte
}

var universityId uuid.UUID
var collegeId uuid.UUID
var subjectId uuid.UUID
var envId uuid.UUID
var problemId uuid.UUID
var syllabusId uuid.UUID
var userId uuid.UUID

func TestNonAuth(t *testing.T) {

	testCaseUni := endpointTest{
		name:   "Test University Endpoint",
		path:   "/university",
		method: "POST",
		body: `{
				"title": "AKTU",
				"description": {
					"dummy": "value"
				}
			}`,
		header: map[string]string{
			"Content-Type": "application/json",
		},
		wantRespCode: 200,
	}

	tc := testCaseUni
	t.Run(tc.name, func(t *testing.T) {
		req, _ := http.NewRequest(tc.method, "http://localhost:8080"+tc.path, strings.NewReader(tc.body))
		for key, value := range tc.header {
			req.Header.Set(key, value)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}

		if res.StatusCode != tc.wantRespCode {
			t.Errorf("Expected status code %d but got %d", tc.wantRespCode, res.StatusCode)
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}
		res.Body.Close()

		var createUniRes models.UniversityAPIResponse
		err = json.Unmarshal(body, &createUniRes)
		if err != nil {
			t.Fatal(err)
		}
		universityId = createUniRes.Message.ID
	})

	testCaseCollege := endpointTest{
		name:   "Test College Endpoint",
		path:   "/college",
		method: "POST",
		body: fmt.Sprintf(`{
			"title": "GL BAJAJ",
			"university_id": "%s",
			"description": "Example description"
		  }`, universityId.String()),
		header: map[string]string{
			"Content-Type": "application/json",
		},
		wantRespCode: 200,
	}

	tc = testCaseCollege
	t.Run(tc.name, func(t *testing.T) {
		req, _ := http.NewRequest(tc.method, "http://localhost:8080"+tc.path, strings.NewReader(tc.body))
		for key, value := range tc.header {
			req.Header.Set(key, value)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}

		if res.StatusCode != tc.wantRespCode {
			t.Errorf("Expected status code %d but got %d", tc.wantRespCode, res.StatusCode)
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}
		res.Body.Close()

		var createCollegeRes models.CollegeAPIResponse
		err = json.Unmarshal(body, &createCollegeRes)
		if err != nil {
			t.Fatal(err)
		}

		collegeId = createCollegeRes.Message.ID
	})

	testCaseSub := endpointTest{
		name:   "Test College Endpoint",
		path:   "/subject",
		method: "POST",
		body: fmt.Sprintf(`{
			"title": "Mathematics2",
			"description": "Study of numbers, quantity, structure, space, and change.",
			"university_id": "%s"
		  }`, universityId.String()),
		header: map[string]string{
			"Content-Type": "application/json",
		},
		wantRespCode: 200,
	}

	tc = testCaseSub
	t.Run(tc.name, func(t *testing.T) {
		req, _ := http.NewRequest(tc.method, "http://localhost:8080"+tc.path, strings.NewReader(tc.body))
		for key, value := range tc.header {
			req.Header.Set(key, value)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}

		if res.StatusCode != tc.wantRespCode {
			t.Errorf("Expected status code %d but got %d", tc.wantRespCode, res.StatusCode)
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}
		res.Body.Close()

		var createSubRes models.SubjectAPIResponse
		err = json.Unmarshal(body, &createSubRes)
		if err != nil {
			t.Fatal(err)
		}
		subjectId = createSubRes.Message.ID
	})

	testCaseSyb := endpointTest{
		name:   "Test College Endpoint",
		path:   "/syllabus",
		method: "POST",
		body: fmt.Sprintf(`{
			"subject_id": "%s",
			"uni_college_id": "%s",
			"syllabus_level": "COLLEGE"
		}`, subjectId.String(), collegeId.String()),
		header: map[string]string{
			"Content-Type": "application/json",
		},
		wantRespCode: 200,
	}

	tc = testCaseSyb
	t.Run(tc.name, func(t *testing.T) {
		req, _ := http.NewRequest(tc.method, "http://localhost:8080"+tc.path, strings.NewReader(tc.body))
		for key, value := range tc.header {
			req.Header.Set(key, value)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}

		if res.StatusCode != tc.wantRespCode {
			t.Errorf("Expected status code %d but got %d", tc.wantRespCode, res.StatusCode)
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}
		res.Body.Close()

		var createSybRes models.SyllabusAPIResponse
		err = json.Unmarshal(body, &createSybRes)
		if err != nil {
			t.Fatal(err)
		}
		syllabusId = createSybRes.Message.ID
	})

	testCaseSignUp := endpointTest{
		name:   "Test Signup Endpoint",
		path:   "/auth/signup",
		method: "POST",
		body: fmt.Sprintf(`{
			"college_id": "%s",
			"email": "ufs61f1@example.com",
			"contact_no": "1234567890",
			"dob": "2006-01-02T15:04:05Z",
			"university_id": "%s",
			"name": "Priyanshu Rajput",
			"user_name":"ivSsf719g1er12",
			"password": "abcd15fs221"
		  }`, collegeId.String(), universityId.String()),
		header: map[string]string{
			"Content-Type": "application/json",
		},
		wantRespCode: 200,
	}
	tc = testCaseSignUp
	t.Run(tc.name, func(t *testing.T) {
		req, _ := http.NewRequest(tc.method, "http://localhost:8080"+tc.path, strings.NewReader(tc.body))
		for key, value := range tc.header {
			req.Header.Set(key, value)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}

		if res.StatusCode != tc.wantRespCode {
			t.Errorf("Expected status code %d but got %d", tc.wantRespCode, res.StatusCode)
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}
		res.Body.Close()

		fmt.Println(string(body))

		var signUpRes models.SignUpAPIResponse
		err = json.Unmarshal(body, &signUpRes)
		if err != nil {
			t.Fatal(err)
		}
		userId = signUpRes.Message.ID
	})

	testCaseEnv := endpointTest{
		name:   "Test Environment Endpoint",
		path:   "/environment",
		method: "POST",
		body: fmt.Sprintf(`{
			"title": "TestLang",
			"created_by": "%s",
			"file":"YXBpVmVyc2lvbjogYmF0Y2gvdjEKa2luZDogSm9iCm1ldGFkYXRhOgogIGFubm90YXRpb25zOgogICAgYmF0Y2gua3ViZXJuZXRlcy5pby9qb2ItdHJhY2tpbmc6ICIiCiAgY3JlYXRpb25UaW1lc3RhbXA6ICIyMDIzLTExLTI3VDE1OjU2OjA2WiIKICBnZW5lcmF0aW9uOiAxCiAgbGFiZWxzOgogICAgYmF0Y2gua3ViZXJuZXRlcy5pby9jb250cm9sbGVyLXVpZDogYzBjOGQ5MWMtNGU5OS00ZGZhLTk2YjYtZWNiZDM0MzA0MTRiCiAgICBiYXRjaC5rdWJlcm5ldGVzLmlvL2pvYi1uYW1lOiBzb2x1dGlvbi0wMDg3ZWM1My0zMzBmLTQ3YzMtOWIwNC02NjU3NWE5ZDBhZjYKICAgIGNvbnRyb2xsZXItdWlkOiBjMGM4ZDkxYy00ZTk5LTRkZmEtOTZiNi1lY2JkMzQzMDQxNGIKICAgIGpvYi1uYW1lOiBzb2x1dGlvbi0wMDg3ZWM1My0zMzBmLTQ3YzMtOWIwNC02NjU3NWE5ZDBhZjYKICBuYW1lOiBzb2x1dGlvbi0wMDg3ZWM1My0zMzBmLTQ3YzMtOWIwNC02NjU3NWE5ZDBhZjYKICBuYW1lc3BhY2U6IHN0b3JhZ2UKICByZXNvdXJjZVZlcnNpb246ICIxMzY4MiIKICB1aWQ6IGMwYzhkOTFjLTRlOTktNGRmYS05NmI2LWVjYmQzNDMwNDE0YgpzcGVjOgogIGJhY2tvZmZMaW1pdDogNAogIGNvbXBsZXRpb25Nb2RlOiBOb25JbmRleGVkCiAgY29tcGxldGlvbnM6IDEKICBwYXJhbGxlbGlzbTogMQogIHNlbGVjdG9yOgogICAgbWF0Y2hMYWJlbHM6CiAgICAgIGJhdGNoLmt1YmVybmV0ZXMuaW8vY29udHJvbGxlci11aWQ6IGMwYzhkOTFjLTRlOTktNGRmYS05NmI2LWVjYmQzNDMwNDE0YgogIHN1c3BlbmQ6IGZhbHNlCiAgdGVtcGxhdGU6CiAgICBtZXRhZGF0YToKICAgICAgY3JlYXRpb25UaW1lc3RhbXA6IG51bGwKICAgICAgbGFiZWxzOgogICAgICAgIGJhdGNoLmt1YmVybmV0ZXMuaW8vY29udHJvbGxlci11aWQ6IGMwYzhkOTFjLTRlOTktNGRmYS05NmI2LWVjYmQzNDMwNDE0YgogICAgICAgIGJhdGNoLmt1YmVybmV0ZXMuaW8vam9iLW5hbWU6IHNvbHV0aW9uLTAwODdlYzUzLTMzMGYtNDdjMy05YjA0LTY2NTc1YTlkMGFmNgogICAgICAgIGNvbnRyb2xsZXItdWlkOiBjMGM4ZDkxYy00ZTk5LTRkZmEtOTZiNi1lY2JkMzQzMDQxNGIKICAgICAgICBqb2ItbmFtZTogc29sdXRpb24tMDA4N2VjNTMtMzMwZi00N2MzLTliMDQtNjY1NzVhOWQwYWY2CiAgICBzcGVjOgogICAgICBjb250YWluZXJzOgogICAgICAtIGFyZ3M6CiAgICAgICAgLSAtLWhlYWRlcgogICAgICAgIC0gJ0NvbnRlbnQtVHlwZTogYXBwbGljYXRpb24vanNvbicKICAgICAgICAtIC0tZGF0YQogICAgICAgIC0gJ3sic2NvcmUiOiA4OSwgInN0YXR1cyI6ICJBY2NlcHRlZCIsICJydW5fdGltZSI6ICIxLjAyIn0nCiAgICAgICAgY29tbWFuZDoKICAgICAgICAtIGN1cmwKICAgICAgICAtIC0tbG9jYXRpb24KICAgICAgICAtIC0tcmVxdWVzdAogICAgICAgIC0gUFVUCiAgICAgICAgLSBodHRwOi8vaG9zdC5kb2NrZXIuaW50ZXJuYWw6ODA4MC9zdWJtaXNzaW9uP2lkPTAwODdlYzUzLTMzMGYtNDdjMy05YjA0LTY2NTc1YTlkMGFmNgogICAgICAgIGltYWdlOiBjdXJsaW1hZ2VzL2N1cmw6bGF0ZXN0CiAgICAgICAgaW1hZ2VQdWxsUG9saWN5OiBBbHdheXMKICAgICAgICBuYW1lOiBwaQogICAgICAgIHJlc291cmNlczoge30KICAgICAgICB0ZXJtaW5hdGlvbk1lc3NhZ2VQYXRoOiAvZGV2L3Rlcm1pbmF0aW9uLWxvZwogICAgICAgIHRlcm1pbmF0aW9uTWVzc2FnZVBvbGljeTogRmlsZQogICAgICAgIHZvbHVtZU1vdW50czoKICAgICAgICAtIG1vdW50UGF0aDogL3BhdGgvdG8vc29sdXRpb24KICAgICAgICAgIG5hbWU6IHNvbHV0aW9uCiAgICAgICAgLSBtb3VudFBhdGg6IC9wYXRoL3RvL3Rlc3QKICAgICAgICAgIG5hbWU6IHRlc3QKICAgICAgZG5zUG9saWN5OiBDbHVzdGVyRmlyc3QKICAgICAgcmVzdGFydFBvbGljeTogTmV2ZXIKICAgICAgc2NoZWR1bGVyTmFtZTogZGVmYXVsdC1zY2hlZHVsZXIKICAgICAgc2VjdXJpdHlDb250ZXh0OiB7fQogICAgICB0ZXJtaW5hdGlvbkdyYWNlUGVyaW9kU2Vjb25kczogMzAKICAgICAgdm9sdW1lczoKICAgICAgLSBjb25maWdNYXA6CiAgICAgICAgICBkZWZhdWx0TW9kZTogNDIwCiAgICAgICAgICBuYW1lOiBzb2x1dGlvbi0wMDg3ZWM1My0zMzBmLTQ3YzMtOWIwNC02NjU3NWE5ZDBhZjYKICAgICAgICBuYW1lOiBzb2x1dGlvbgogICAgICAtIGNvbmZpZ01hcDoKICAgICAgICAgIGRlZmF1bHRNb2RlOiA0MjAKICAgICAgICAgIG5hbWU6IHRlc3QtZmlsZS0yM2Q1OTI5ZS1jOWRmLTQ2YjktYjRmZC0wMTZkOWQ4Nzc4NDIKICAgICAgICBuYW1lOiB0ZXN0CnN0YXR1czoKICBjb21wbGV0aW9uVGltZTogIjIwMjMtMTEtMjdUMTU6NTY6MTJaIgogIGNvbmRpdGlvbnM6CiAgLSBsYXN0UHJvYmVUaW1lOiAiMjAyMy0xMS0yN1QxNTo1NjoxMloiCiAgICBsYXN0VHJhbnNpdGlvblRpbWU6ICIyMDIzLTExLTI3VDE1OjU2OjEyWiIKICAgIHN0YXR1czogIlRydWUiCiAgICB0eXBlOiBDb21wbGV0ZQogIHJlYWR5OiAwCiAgc3RhcnRUaW1lOiAiMjAyMy0xMS0yN1QxNTo1NjowNloiCiAgc3VjY2VlZGVkOiAxCiAgdW5jb3VudGVkVGVybWluYXRlZFBvZHM6IHt9Cg=="
		}`, userId.String()),
		header: map[string]string{
			"Content-Type": "application/json",
		},
		wantRespCode: 200,
	}
	tc = testCaseEnv
	t.Run(tc.name, func(t *testing.T) {
		req, _ := http.NewRequest(tc.method, "http://localhost:8080"+tc.path, strings.NewReader(tc.body))
		for key, value := range tc.header {
			req.Header.Set(key, value)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}

		if res.StatusCode != tc.wantRespCode {
			t.Errorf("Expected status code %d but got %d", tc.wantRespCode, res.StatusCode)
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}
		res.Body.Close()
		fmt.Println(string(body))
		var createEnvRes models.EnvironmentAPIResponse
		err = json.Unmarshal(body, &createEnvRes)
		if err != nil {
			t.Fatal(err)
		}
		envId = createEnvRes.Message.ID
	})

	testCaseProb := endpointTest{
		name:   "Test Environment Endpoint",
		path:   "/problem",
		method: "POST",
		body: fmt.Sprintf(`{
			"title": "Implement CPU Scheduling Policy SJF",
			"created_by": "%s",
			"difficulty": "MEDIUM",
			"syllabus_id": "%s",
			"environment": [
			  {
				"language": "Go",
				"id": "%s"
			  }
			],
			"problem_file":"DQogICAgPGgyPlByb2JsZW0gU3RhdGVtZW50OjwvaDI+DQogICAgPHA+DQogICAgICAgIERlc2lnbiBhbmQgSW1wbGVtZW50YXRpb24gb2YgUGF5cm9sbCBQcm9jZXNzaW5nIFN5c3RlbQ0KICAgIDwvcD4NCg0KICAgIDxwPg0KICAgICAgICBJbiB0b2RheSdzIGR5bmFtaWMgYnVzaW5lc3MgZW52aXJvbm1lbnQsIGVmZmljaWVudCBtYW5hZ2VtZW50IG9mIGVtcGxveWVlIHBheXJvbGwgaXMgY3J1Y2lhbCBmb3Igb3JnYW5pemF0aW9uYWwgc3VjY2Vzcy4gWW91ciB0YXNrIGlzIHRvIGRlc2lnbiBhbmQgaW1wbGVtZW50IGEgUGF5cm9sbCBQcm9jZXNzaW5nIFN5c3RlbSBmb3IgYSBjb21wYW55Lg0KICAgIDwvcD4NCg0KICAgIDxwPg0KICAgICAgICBUaGUgc3lzdGVtIHNob3VsZCBiZSBjYXBhYmxlIG9mIGhhbmRsaW5nIGVtcGxveWVlIGluZm9ybWF0aW9uLCBkZXBhcnRtZW50IGRldGFpbHMsIHBvc2l0aW9uIGRldGFpbHMsIGFsbG93YW5jZXMsIGRlZHVjdGlvbnMsIGJhc2UgcGF5LCBhbmQgcGF5cm9sbCBwcm9jZXNzaW5nLg0KICAgIDwvcD4NCg0KICAgIDxoMz5UYWJsZSAiRW1wbG95ZWUiPC9oMz4NCiAgICA8cHJlPg0KICAgICAgICAiRW1wbG95ZWVJRCIgSU5UIFtwa10NCiAgICAgICAgIkZpcnN0TmFtZSIgVkFSQ0hBUigyNTUpDQogICAgICAgICJMYXN0TmFtZSIgVkFSQ0hBUigyNTUpDQogICAgICAgICJEYXRlT2ZCaXJ0aCIgREFURQ0KICAgICAgICAiR2VuZGVyIiBDSEFSKDEpDQogICAgICAgICJDb250YWN0TnVtYmVyIiBWQVJDSEFSKDIwKSBbdW5pcXVlXQ0KICAgICAgICAiRW1haWwiIFZBUkNIQVIoMjU1KSBbdW5pcXVlXQ0KICAgICAgICAiQWRkcmVzcyIgVkFSQ0hBUig1MDApDQogICAgICAgICJCYW5rQWNjb3VudE51bWJlciIgVkFSQ0hBUig1MCkgW3VuaXF1ZV0NCiAgICAgICAgIlRheElEIiBWQVJDSEFSKDUwKSBbdW5pcXVlXQ0KICAgICAgICAiUG9zaXRpb24iIFZBUkNIQVIoMTAwKQ0KICAgICAgICAiRGVwYXJ0bWVudElEIiBJTlQNCiAgICA8L3ByZT4NCg0KICAgIDxoMz5UYWJsZSAiRGVwYXJ0bWVudCI8L2gzPg0KICAgIDxwcmU+DQogICAgICAgICJEZXBhcnRtZW50SUQiIElOVCBbcGtdDQogICAgICAgICJEZXBhcnRtZW50TmFtZSIgVkFSQ0hBUigxMDApIFt1bmlxdWVdDQogICAgPC9wcmU+DQoNCiAgICA8aDQ+SW5zZXJ0ZWQgRGF0YSBpbiBEZXBhcnRtZW50IFRhYmxlOjwvaDQ+DQogICAgPHByZT4NCiAgICAgICAgKERlcGFydG1lbnRJRCwgRGVwYXJ0bWVudE5hbWUpIA0KICAgICAgICAoMTAxLCAnRW5naW5lZXJpbmcnKSwNCiAgICAgICAgKDEwMiwgJ1NhbGVzJyksDQogICAgICAgICgxMDMsICdIdW1hbiBSZXNvdXJjZXMnKTsNCiAgICA8L3ByZT4NCg0KICAgIDxoND5JbnNlcnRlZCBEYXRhIGluIEVtcGxveWVlIFRhYmxlOjwvaDQ+DQogICAgPHByZT4NCiAgICAgICAoRW1wbG95ZWVJRCwgRmlyc3ROYW1lLCBMYXN0TmFtZSwgRGF0ZU9mQmlydGgsIEdlbmRlciwgQ29udGFjdE51bWJlciwgRW1haWwsIEFkZHJlc3MsIEJhbmtBY2NvdW50TnVtYmVyLCBUYXhJRCwgUG9zaXRpb24sIERlcGFydG1lbnRJRCkgDQogICAgICAgICgxLCAnSm9obicsICdEb2UnLCAnMTk5MC0wMS0xNScsICdNJywgJzEyMzQ1Njc4OTAnLCAnam9obi5kb2VAZXhhbXBsZS5jb20nLCAnMTIzIE1haW4gU3QsIENpdHksIENvdW50cnknLCAnOTg3NjU0MzIxJywgJ1RBWDEyMycsICdTb2Z0d2FyZSBFbmdpbmVlcicsIDEwMSksDQogICAgICAgICgyLCAnSmFuZScsICdTbWl0aCcsICcxOTg1LTA1LTIwJywgJ0YnLCAnOTg3NjU0MzIxMCcsICdqYW5lLnNtaXRoQGV4YW1wbGUuY29tJywgJzQ1NiBPYWsgU3QsIENpdHksIENvdW50cnknLCAnMTIzNDU2Nzg5JywgJ1RBWDQ1NicsICdTYWxlcyBNYW5hZ2VyJywgMTAyKSwNCiAgICAgICAgKDMsICdCb2InLCAnSm9obnNvbicsICcxOTkyLTA5LTMwJywgJ00nLCAnNTU1MjIyMTExMScsICdib2Iuam9obnNvbkBleGFtcGxlLmNvbScsICc3ODkgUGluZSBTdCwgQ2l0eSwgQ291bnRyeScsICc0NTY3ODkwMTInLCAnVEFYNzg5JywgJ0hSIFNwZWNpYWxpc3QnLCAxMDMpLA0KICAgICAgICAoNCwgJ0FsaWNlJywgJ1dpbGxpYW1zJywgJzE5ODgtMDMtMTAnLCAnRicsICczMzM0NDQ1NTU1JywgJ2FsaWNlLndpbGxpYW1zQGV4YW1wbGUuY29tJywgJzc4OSBFbG0gU3QsIENpdHksIENvdW50cnknLCAnMjM0NTY3ODkwJywgJ1RBWDIzNCcsICdTb2Z0d2FyZSBEZXZlbG9wZXInLCAxMDEpLA0KICAgICAgICAoNSwgJ01pY2hhZWwnLCAnQnJvd24nLCAnMTk5NS0xMi0yMicsICdNJywgJzc3Nzg4ODk5OTknLCAnbWljaGFlbC5icm93bkBleGFtcGxlLmNvbScsICczMjEgQmlyY2ggU3QsIENpdHksIENvdW50cnknLCAnNTY3ODkwMTIzJywgJ1RBWDU2NycsICdTYWxlcyBSZXByZXNlbnRhdGl2ZScsIDEwMiksDQogICAgICAgICg2LCAnRW1pbHknLCAnTWlsbGVyJywgJzE5ODItMDYtMDUnLCAnRicsICc4ODg5OTkwMDAwJywgJ2VtaWx5Lm1pbGxlckBleGFtcGxlLmNvbScsICc2NTQgUGluZSBTdCwgQ2l0eSwgQ291bnRyeScsICc2Nzg5MDEyMzQnLCAnVEFYNjc4JywgJ0hSIE1hbmFnZXInLCAxMDMpOw0KICAgIDwvcHJlPg0KDQogICAgPHByZT4NCiAgICAgICAgUmVmOiJEZXBhcnRtZW50Ii4iRGVwYXJ0bWVudElEIiA8ICJFbXBsb3llZSIuIkRlcGFydG1lbnRJRCINCiAgICA8L3ByZT4NCg0KICAgIDxoMj5UYXNrIDE6PC9oMj4NCg0KICAgIDxvbD4NCiAgICA8bGk+RW1wbG95ZWUgYW5kIERlcGFydG1lbnQgVGFibGVzIGFyZSBhbHJlYWR5IGNyZWF0ZWQsIHdyaXRlIFNRTCBRdWVyeSB0byBjcmVhdGUgU2FsYXJ5IGFuZCBUaW1lc2hlZXQgdGFibGVzLg0KDQogICAgPHA+VGFibGUgIlNhbGFyeSI8L3A+DQogICAgPHByZT4NCiAgICAgICAgIlNhbGFyeUlEIiBJTlQgW3BrXQ0KICAgICAgICAiRW1wbG95ZWVJRCIgSU5UIFtmb3JlaWduOiA+IEVtcGxveWVlLkVtcGxveWVlSURdDQogICAgICAgICJFZmZlY3RpdmVEYXRlIiBEQVRFDQogICAgICAgICJCYXNlU2FsYXJ5IiBERUNJTUFMKDEwLDIpDQogICAgICAgICJCb251cyIgREVDSU1BTCgxMCwyKQ0KICAgICAgICAiRGVkdWN0aW9ucyIgREVDSU1BTCgxMCwyKQ0KICAgICAgICAiTmV0U2FsYXJ5IiBERUNJTUFMKDEwLDIpDQogICAgPC9wcmU+DQoNCiAgICA8cD5UYWJsZSAiVGltZVNoZWV0IjwvcD4NCiAgICA8cHJlPg0KICAgICAgICAiVGltZVNoZWV0SUQiIElOVCBbcGtdDQogICAgICAgICJFbXBsb3llZUlEIiBJTlQgW2ZvcmVpZ246ID4gRW1wbG95ZWUuRW1wbG95ZWVJRF0NCiAgICAgICAgIldvcmtEYXRlIiBEQVRFDQogICAgICAgICJIb3Vyc1dvcmtlZCIgREVDSU1BTCg1LDIpDQogICAgICAgICJPdmVydGltZUhvdXJzIiBERUNJTUFMKDUsMikNCiAgICA8L3ByZT4NCiAgICA8L2xpPg0KDQogICAgPGxpPldyaXRlIFF1ZXJ5IHRvIGluc2VydCBmb2xsb3dpbmcgZGF0YSBpbnRvIFNhbGFyeSBhbmQgVGltZVNoZWV0IHRhYmxlDQoNCiAgICA8aDU+IlNhbGFyeSIgdGFibGUgLT4gKCJFbXBsb3llZUlEIiwgIkVmZmVjdGl2ZURhdGUiLCAiQmFzZVNhbGFyeSIsICJCb251cyIsICJEZWR1Y3Rpb25zIiwgIk5ldFNhbGFyeSIpPC9oNT4NCiAgICA8cHJlPg0KICAgICAgICBWQUxVRVMgLT4gICgxLCAnMjAyMy0wMS0wMScsIDgwMDAwLjAwLCA1MDAwLjAwLCAyMDAwLjAwLCA4MzAwMC4wMCksDQogICAgICAgICAgICAgICAgKDEsICcyMDIzLTA3LTAxJywgODUwMDAuMDAsIDYwMDAuMDAsIDI1MDAuMDAsIDg2NTAwLjAwKTsNCiAgICA8L3ByZT4NCg0KICAgIDxoNT4iVGltZVNoZWV0IiB0YWJsZSAtPiAoIkVtcGxveWVlSUQiLCAiV29ya0RhdGUiLCAiSG91cnNXb3JrZWQiLCAiT3ZlcnRpbWVIb3VycyIpPC9oNT4NCiAgICA8cHJlPg0KICAgICAgICBWQUxVRVMgLT4gKDEsICcyMDIzLTEwLTAxJywgOC4wMCwgMi4wMCksDQogICAgICAgICAgICAgICAgKDEsICcyMDIzLTEwLTAyJywgNy41MCwgMS41MCk7DQogICAgPC9wcmU+DQogICAgPC9saT4NCiAgICA8L29sPg0KICAgIDxoMj5UYXNrIDI6PC9oMj4NCg0KICAgIDxvbD4NCiAgICA8bGk+V3JpdGUgYSBxdWVyeSB0byBnZXQgYWxsIG1hbGUgZW1wbG95ZWVzJyBlbWFpbCBzb3J0ZWQgYnkgZGVwYXJ0bWVudCBpZC48L2xpPg0KICAgIDxsaT5Xcml0ZSBhIHF1ZXJ5IHRvIGdldCB0aGUgb2xkZXN0IGZlbWFsZSBlbXBsb3llZSBhbGwgZGV0YWlscy48L2xpPg0KICAgIDwvb2w+DQoNCiAgICA8aDI+VGFzayAzOjwvaDI+DQoNCiAgICA8b2w+DQogICAgPGxpPldyaXRlIGEgcXVlcnkgdG8gZ2V0IHRvdGFsIG1vbmV5IGVhcm5lZCBieSBlbXBsb3llZSAxIG9uICcyMDIzLTAxLTAxJyBhbmQgJzIwMjMtMDctMDEnPC9saT4NCiAgICA8bGk+VXBkYXRlIGVtcGxveWVlIDEgc2FsYXJ5IGJvbnVzIGFuZCBuZXQgc2FsYXJ5IHRvIDYwMDAsIDg0MDAwIG9uICcyMDIzLTAxLTAxJzwvbGk+DQogICAgPGxpPlJlbW92ZSBUaW1lU2hlZXQgZW50cnkgZm9yIGRhdGUgJzIwMjMtMTAtMDInPC9saT4NCiAgICA8L29sPg0K",
			"test_files": [
			  {
				"language": "MYSQL",
				"title" : "mysql",
				"init_code": "4oCUVGFzayAxOiBTUUwgUXVlcmllcyBHb2VzIEhlcmUuLgoK4oCTVGFzayAyOiBTUUwgUXVlcmllcyBHb2VzIEhlcmUuLgoK4oCTVGFzayAzOiBTUUwgUXVlcmllcyBHb2VzIEhlcmUuLgo=",
				"file": "aW1wb3J0IHN5cwppbXBvcnQganNvbgoKZGVmIHJ1bl90ZXN0KHRlc3RfaW5wdXQsIGV4cGVjdGVkX291dHB1dCwgZnVuY3Rpb25fbmFtZSk6CiAgICB0cnk6CiAgICAgICAgIyBFeGVjdXRlIHRoZSBzb2x1dGlvbiBjb2RlIGluIHRoZSBnbG9iYWwgc2NvcGUKICAgICAgICBnbG9iYWxzX2RpY3QgPSB7fQogICAgICAgIGV4ZWMoc29sdXRpb25fY29kZSwgZ2xvYmFsc19kaWN0KQogICAgICAgIGlmIGZ1bmN0aW9uX25hbWUgbm90IGluIGdsb2JhbHNfZGljdCBvciBub3QgY2FsbGFibGUoZ2xvYmFsc19kaWN0W2Z1bmN0aW9uX25hbWVdKToKICAgICAgICAgICAgcmV0dXJuIHsidGVzdF9jYXNlIjogdGVzdF9pbnB1dCwgImVycm9yIjogIkVycm9yOiBUaGUgZnVuY3Rpb24ge30gaXMgbm90IGRlZmluZWQuIi5mb3JtYXQoZnVuY3Rpb25fbmFtZSksICJyZXN1bHQiOiAiRmFpbGVkIn0KCiAgICAgICAgIyBUZXN0IHRoZSBmdW5jdGlvbiB3aXRoIHRoZSBwcm92aWRlZCBpbnB1dAogICAgICAgIHJlc3VsdCA9IGdsb2JhbHNfZGljdFtmdW5jdGlvbl9uYW1lXSh0ZXN0X2lucHV0KQoKICAgICAgICAjIENoZWNrIGlmIHRoZSByZXN1bHQgbWF0Y2hlcyB0aGUgZXhwZWN0ZWQgb3V0cHV0CiAgICAgICAgaWYgcmVzdWx0ICE9IGV4cGVjdGVkX291dHB1dDoKICAgICAgICAgICAgcmV0dXJuIHsidGVzdF9jYXNlIjogdGVzdF9pbnB1dCwgImVycm9yIjogTm9uZSwgInJlc3VsdCI6ICJGYWlsZWQiLCAiZGV0YWlscyI6IHsiZXhwZWN0ZWRfb3V0cHV0IjogZXhwZWN0ZWRfb3V0cHV0LCAiYWN0dWFsX291dHB1dCI6IHJlc3VsdH19CgogICAgICAgICMgSWYgdGhlIHJlc3VsdCBtYXRjaGVzIHRoZSBleHBlY3RlZCBvdXRwdXQsIHJldHVybiBzdWNjZXNzCiAgICAgICAgcmV0dXJuIHsidGVzdF9jYXNlIjogdGVzdF9pbnB1dCwgImVycm9yIjogTm9uZSwgInJlc3VsdCI6ICJQYXNzZWQifQoKICAgIGV4Y2VwdCBFeGNlcHRpb24gYXMgZToKICAgICAgICByZXR1cm4geyJ0ZXN0X2Nhc2UiOiB0ZXN0X2lucHV0LCAiZXJyb3IiOiAiRXJyb3I6IHt9Ii5mb3JtYXQoc3RyKGUpKSwgInJlc3VsdCI6ICJGYWlsZWQifQoKZGVmIGNoZWNrX2ZpYm9uYWNjaV9zb2x1dGlvbihzb2x1dGlvbl9maWxlX3BhdGgsIHRlc3RfY2FzZXMpOgogICAgdHJ5OgogICAgICAgIHdpdGggb3Blbihzb2x1dGlvbl9maWxlX3BhdGgsICdyJykgYXMgZmlsZToKICAgICAgICAgICAgZ2xvYmFsIHNvbHV0aW9uX2NvZGUKICAgICAgICAgICAgc29sdXRpb25fY29kZSA9IGZpbGUucmVhZCgpCgogICAgICAgIHJlc3VsdHMgPSBbXQogICAgICAgIGZvciB0ZXN0X2Nhc2UgaW4gdGVzdF9jYXNlczoKICAgICAgICAgICAgcmVzdWx0ID0gcnVuX3Rlc3QodGVzdF9jYXNlWyJpbnB1dCJdLCB0ZXN0X2Nhc2VbIm91dHB1dCJdLCB0ZXN0X2Nhc2VbImZ1bmN0aW9uX25hbWUiXSkKICAgICAgICAgICAgcmVzdWx0cy5hcHBlbmQocmVzdWx0KQoKICAgICAgICByZXR1cm4ganNvbi5kdW1wcyhyZXN1bHRzLCBpbmRlbnQ9MikKCiAgICBleGNlcHQgRXhjZXB0aW9uIGFzIGU6CiAgICAgICAgcmV0dXJuIGpzb24uZHVtcHMoW3siZXJyb3IiOiAiRXJyb3I6IHt9Ii5mb3JtYXQoc3RyKGUpKX1dLCBpbmRlbnQ9MikKCmlmIF9fbmFtZV9fID09ICJfX21haW5fXyI6CiAgICBpZiBsZW4oc3lzLmFyZ3YpICE9IDI6CiAgICAgICAgcHJpbnQoIlVzYWdlOiBweXRob24gY2hlY2tfc29sdXRpb24ucHkgPHNvbHV0aW9uX2ZpbGVfcGF0aD4iKQogICAgICAgIHN5cy5leGl0KDEpCgogICAgIyBEZWZpbmUgMTAgdGVzdCBjYXNlcyB3aXRoIGlucHV0LCBleHBlY3RlZCBvdXRwdXQsIGFuZCBmdW5jdGlvbiBuYW1lCiAgICB0ZXN0X2Nhc2VzID0gWwogICAgICAgIHsiaW5wdXQiOiAwLCAib3V0cHV0IjogW10sICJmdW5jdGlvbl9uYW1lIjogImdlbmVyYXRlX2ZpYm9uYWNjaSJ9LAogICAgICAgIHsiaW5wdXQiOiAxLCAib3V0cHV0IjogWzBdLCAiZnVuY3Rpb25fbmFtZSI6ICJnZW5lcmF0ZV9maWJvbmFjY2kifSwKICAgICAgICB7ImlucHV0IjogMiwgIm91dHB1dCI6IFswLCAxXSwgImZ1bmN0aW9uX25hbWUiOiAiZ2VuZXJhdGVfZmlib25hY2NpIn0sCiAgICAgICAgeyJpbnB1dCI6IDUsICJvdXRwdXQiOiBbMCwgMSwgMSwgMiwgM10sICJmdW5jdGlvbl9uYW1lIjogImdlbmVyYXRlX2ZpYm9uYWNjaSJ9LAogICAgICAgIHsiaW5wdXQiOiAxMCwgIm91dHB1dCI6IFswLCAxLCAxLCAyLCAzLCA1LCA4LCAxMywgMjEsIDM0XSwgImZ1bmN0aW9uX25hbWUiOiAiZ2VuZXJhdGVfZmlib25hY2NpIn0sCiAgICAgICAgeyJpbnB1dCI6IDE1LCAib3V0cHV0IjogWzAsIDEsIDEsIDIsIDMsIDUsIDgsIDEzLCAyMSwgMzQsIDU1LCA4OSwgMTQ0LCAyMzMsIDM3N10sICJmdW5jdGlvbl9uYW1lIjogImdlbmVyYXRlX2ZpYm9uYWNjaSJ9LAogICAgICAgIHsiaW5wdXQiOiAyMCwgIm91dHB1dCI6IFswLCAxLCAxLCAyLCAzLCA1LCA4LCAxMywgMjEsIDM0LCA1NSwgODksIDE0NCwgMjMzLCAzNzcsIDYxMCwgOTg3LCAxNTk3LCAyNTg0LCA0MTgxXSwgImZ1bmN0aW9uX25hbWUiOiAiZ2VuZXJhdGVfZmlib25hY2NpIn0sCiAgICAgICAgeyJpbnB1dCI6IDI1LCAib3V0cHV0IjogWzAsIDEsIDEsIDIsIDMsIDUsIDgsIDEzLCAyMSwgMzQsIDU1LCA4OSwgMTQ0LCAyMzMsIDM3NywgNjEwLCA5ODcsIDE1OTcsIDI1ODQsIDQxODEsIDY3NjUsIDEwOTQ2LCAxNzcxMSwgMjg2NTcsIDQ2MzY4XSwgImZ1bmN0aW9uX25hbWUiOiAiZ2VuZXJhdGVfZmlib25hY2NpIn0sCiAgICAgICAgeyJpbnB1dCI6IDMwLCAib3V0cHV0IjogWzAsIDEsIDEsIDIsIDMsIDUsIDgsIDEzLCAyMSwgMzQsIDU1LCA4OSwgMTQ0LCAyMzMsIDM3NywgNjEwLCA5ODcsIDE1OTcsIDI1ODQsIDQxODEsIDY3NjUsIDEwOTQ2LCAxNzcxMSwgMjg2NTcsIDQ2MzY4LCA3NTAyNSwgMTIxMzkzLCAxOTY0MTgsIDMxNzgxMSwgNTE0MjI5XSwgImZ1bmN0aW9uX25hbWUiOiAiZ2VuZXJhdGVfZmlib25hY2NpIn0sCiAgICBdCgogICAgc29sdXRpb25fZmlsZV9wYXRoID0gc3lzLmFyZ3ZbMV0KICAgIHJlc3VsdCA9IGNoZWNrX2ZpYm9uYWNjaV9zb2x1dGlvbihzb2x1dGlvbl9maWxlX3BhdGgsIHRlc3RfY2FzZXMpCiAgICBwcmludChyZXN1bHQpCg=="
			  }
			]
		
		}`, userId, syllabusId, envId),
		header: map[string]string{
			"Content-Type": "application/json",
		},
		wantRespCode: 200,
	}
	tc = testCaseProb
	t.Run(tc.name, func(t *testing.T) {
		req, _ := http.NewRequest(tc.method, "http://localhost:8080"+tc.path, strings.NewReader(tc.body))
		for key, value := range tc.header {
			req.Header.Set(key, value)
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}

		if res.StatusCode != tc.wantRespCode {
			t.Errorf("Expected status code %d but got %d", tc.wantRespCode, res.StatusCode)
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}
		res.Body.Close()
		fmt.Println(string(body))
		var createProbRes models.EnvironmentAPIResponse
		err = json.Unmarshal(body, &createProbRes)
		if err != nil {
			t.Fatal(err)
		}
		problemId = createProbRes.Message.ID
	})
}
