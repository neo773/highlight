defmodule Highlight.Config do
  @moduledoc """
  Stores configuration variables used for Highlight.io SDK.
  """

  @doc """
  Returns the Highlight project ID. Set it in `config/config.exs`:

      config :highlight, project_id: "YOUR_PROJECT_ID"
  """
  def project_id, do: from_env(:highlight, :project_id)

  @doc """
  Returns the Highlight API endpoint. Set it in `config/config.exs`:

      config :highlight, api_endpoint: "https://api.highlight.io"
  """
  def api_endpoint, do: from_env(:highlight, :api_endpoint, "https://api.highlight.io")

  @doc """
  Returns the environment setting. Set it in `config/config.exs`:

      config :highlight, environment: "production"
  """
  def environment, do: from_env(:highlight, :environment, "production")

  @doc """
  A helper function to read configuration values, supporting `{:system, "VAR"}` tuples.
  """
  def from_env(app, key, default \\ nil) do
    app
    |> Application.get_env(key, default)
    |> read_from_system(default)
  end

  defp read_from_system({:system, env_var}, default), do: System.get_env(env_var) || default
  defp read_from_system(value, _default), do: value
end
