package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"bytes"

	"github.com/docker-engine-api/mcp-server/config"
	"github.com/docker-engine-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func ContainercreateHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["name"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("name=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		// Create properly typed request body using the generated schema
		var requestBody interface{}
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/containers/create%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// No authentication required for this endpoint
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result map[string]interface{}
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateContainercreateTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_containers_create",
		mcp.WithDescription("Create a container"),
		mcp.WithString("name", mcp.Description("Assign the specified name to the container. Must match `/?[a-zA-Z0-9_-]+`.")),
		mcp.WithString("Image", mcp.Description("Input parameter: The name of the image to use when creating the container")),
		mcp.WithString("StopSignal", mcp.Description("Input parameter: Signal to stop a container as a string or unsigned integer.")),
		mcp.WithString("User", mcp.Description("Input parameter: The user that commands are run as inside the container.")),
		mcp.WithArray("OnBuild", mcp.Description("Input parameter: `ONBUILD` metadata that were defined in the image's `Dockerfile`.")),
		mcp.WithObject("Volumes", mcp.Description("Input parameter: An object mapping mount point paths inside the container to empty objects.")),
		mcp.WithObject("Healthcheck", mcp.Description("Input parameter: A test to perform to check that the container is healthy.")),
		mcp.WithBoolean("Tty", mcp.Description("Input parameter: Attach standard streams to a TTY, including `stdin` if it is not closed.")),
		mcp.WithString("Domainname", mcp.Description("Input parameter: The domain name to use for the container.")),
		mcp.WithBoolean("ArgsEscaped", mcp.Description("Input parameter: Command is already escaped (Windows only)")),
		mcp.WithString("Cmd", mcp.Description("Input parameter: Command to run specified as a string or an array of strings.")),
		mcp.WithBoolean("OpenStdin", mcp.Description("Input parameter: Open `stdin`")),
		mcp.WithArray("Env", mcp.Description("Input parameter: A list of environment variables to set inside the container in the form `[\"VAR=value\", ...]`. A variable without `=` is removed from the environment, rather than to have an empty value.\n")),
		mcp.WithBoolean("AttachStdout", mcp.Description("Input parameter: Whether to attach to `stdout`.")),
		mcp.WithNumber("StopTimeout", mcp.Description("Input parameter: Timeout to stop a container in seconds.")),
		mcp.WithString("Entrypoint", mcp.Description("Input parameter: The entry point for the container as a string or an array of strings.\n\nIf the array consists of exactly one empty string (`[\"\"]`) then the entry point is reset to system default (i.e., the entry point used by docker when there is no `ENTRYPOINT` instruction in the `Dockerfile`).\n")),
		mcp.WithString("MacAddress", mcp.Description("Input parameter: MAC address of the container.")),
		mcp.WithBoolean("NetworkDisabled", mcp.Description("Input parameter: Disable networking for the container.")),
		mcp.WithObject("ExposedPorts", mcp.Description("Input parameter: An object mapping ports to an empty object in the form:\n\n`{\"<port>/<tcp|udp>\": {}}`\n")),
		mcp.WithArray("Shell", mcp.Description("Input parameter: Shell for when `RUN`, `CMD`, and `ENTRYPOINT` uses a shell.")),
		mcp.WithString("WorkingDir", mcp.Description("Input parameter: The working directory for commands to run in.")),
		mcp.WithString("Hostname", mcp.Description("Input parameter: The hostname to use for the container, as a valid RFC 1123 hostname.")),
		mcp.WithBoolean("StdinOnce", mcp.Description("Input parameter: Close `stdin` after one attached client disconnects")),
		mcp.WithObject("Labels", mcp.Description("Input parameter: User-defined key/value metadata.")),
		mcp.WithBoolean("AttachStderr", mcp.Description("Input parameter: Whether to attach to `stderr`.")),
		mcp.WithBoolean("AttachStdin", mcp.Description("Input parameter: Whether to attach to `stdin`.")),
		mcp.WithObject("NetworkingConfig", mcp.Description("Input parameter: This container's networking configuration.")),
		mcp.WithObject("HostConfig", mcp.Description("Input parameter: Container configuration that depends on the host we are running on")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    ContainercreateHandler(cfg),
	}
}
