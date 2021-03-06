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

// refreshes a table, this is done by calling StorageServiceMBean#loadNewSSTables 
type RefreshOperationRequest struct {
	Type_ string `json:"type"`
	// keyspace to refresh 
	Keyspace string `json:"keyspace"`
	// table to refresh 
	Table string `json:"table"`
}
