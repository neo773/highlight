defmodule HighlightElixirTest do
  use ExUnit.Case
  doctest HighlightElixir

  test "greets the world" do
    assert HighlightElixir.hello() == :world
  end
end
