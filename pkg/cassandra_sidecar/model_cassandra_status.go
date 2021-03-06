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

type CassandraStatus struct {
	// State of a Cassandra node as reported from StorageServiceMBean#getOperationMode
	NodeState string `json:"nodeState,omitempty"`
}
