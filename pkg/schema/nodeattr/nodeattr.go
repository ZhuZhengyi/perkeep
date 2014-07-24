/*
Copyright 2014 The Camlistore Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package nodeattr contains constants for permanode attribute names.
package nodeattr

const (
	// DateCreated is http://schema.org/dateCreated: The date on
	// which an item was created, in RFC 3339 format (with
	// Camlistore's addition that zone -00:01 means localtime:
	// unknown timezone).
	DateCreated = "dateCreated"
)