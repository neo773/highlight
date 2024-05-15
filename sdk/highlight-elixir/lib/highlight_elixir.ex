defmodule Highlight.Tracer do
  @moduledoc """
  A wrapper around OpenTelemetry for easy integration with Highlight.io.
  """

  require OpenTelemetry.Tracer, as: Tracer
  require OpenTelemetry.Span

  @doc """
  Starts a new span with the given name and attributes.
  """
  def with_span(name, attributes \\ %{}, do: block) do
    Tracer.with_span name, attributes do
      block.()
    end
  end

  @doc """
  Adds an event to the current span.
  """
  def add_event(name, attributes \\ %{}) do
    Tracer.add_event(name, attributes)
  end

  @doc """
  Sets attributes for the current span.
  """
  def set_attributes(attributes) when is_map(attributes) do
    Tracer.set_attributes(attributes)
  end

  @doc """
  Ends the current span.
  """
  def end_span do
    OpenTelemetry.Span.end_span()
  end
end
