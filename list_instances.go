package main
// package listInstances

import (
        "context"
        "fmt"
        "os"
        // "errors"
        // "io"

        compute "cloud.google.com/go/compute/apiv1"
        "google.golang.org/api/iterator"
        computepb "google.golang.org/genproto/googleapis/cloud/compute/v1"
)

func main(){
  PROJECT_ID := os.Getenv("PROJECT_ID")
  ZONE := os.Getenv("ZONE")

  listInstances(PROJECT_ID, ZONE)
  // err := errors.New(list_instances)
  // if list_instances_err != nil {
  //   fmt.Println("an error")
  //   fmt.Println(list_instances_err.Error())
  // }
}

// listInstances prints a list of instances created in given project in given zone.
// func listInstances(w io.Writer, projectID, zone string) error {
func listInstances(projectID, zone string) error {
        // projectID := "your_project_id"
        // zone := "europe-central2-b"
        ctx := context.Background()
        instancesClient, err := compute.NewInstancesRESTClient(ctx)
        if err != nil {
                return fmt.Errorf("NewInstancesRESTClient: %v", err)
        }
        defer instancesClient.Close()

        req := &computepb.ListInstancesRequest{
                Project: projectID,
                Zone:    zone,
        }

        it := instancesClient.List(ctx, req)
        // fmt.Println("Instances found in zone", it[0])

        // fmt.Println("hello")

        for {
                instance, err := it.Next()
                if err == iterator.Done {
                        break
                }
                if err != nil {
                        // fmt.Println("an error")
                        // fmt.Println(err.Error())
                        // return err
                        // fmt.Println("Hello: %v", err)
                        fmt.Errorf("error computing vector: %w", err)
                        return err
                }
                // fmt.Fprintf(w, "- %s %s\n", instance.GetName(), instance.GetMachineType())
                // fmt.Fprintf("- %s %s\n", instance.GetName(), instance.GetMachineType())
                fmt.Println("Instance: " + instance.GetName())
        }
        return nil
}
