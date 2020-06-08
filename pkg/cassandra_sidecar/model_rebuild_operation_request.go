/*
 * Instaclustr Cassandra Sidecar
 *
 * REST API for Cassandra Sidecar from Instaclustr
 *
 * API version: 2.0.0
 * Contact: support@instaclustr.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package cassandra_sidecar

// rebuilds data by streaming from other nodes, 
type RebuildOperationRequest struct {
	Type_ string `json:"type"`
	// specific keyspace to rebuild, if not specified, all keyspaces are rebuilt 
	Keyspace string `json:"keyspace,omitempty"`
	// name of DC from which to select sources for streaming, by default, pick any DC 
	SourceDC string `json:"sourceDC,omitempty"`
	// rebuild specific token ranges 
	SpecificTokens []TokenRange `json:"specificTokens,omitempty"`
	// specify hosts that this node should stream from when specificTokens are used 
	SpecificSources []string `json:"specificSources,omitempty"`
}