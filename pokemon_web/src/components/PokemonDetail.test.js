import { render, screen } from "@testing-library/react";
import PokemonDetail from "./PokemonDetail";

test("render loading text", () => {
  render(
    <PokemonDetail image="https://unpkg.com/pokeapi-sprites@2.0.2/sprites/pokemon/other/dream-world/1.svg" />
  );
  const imageElement = screen.getByText(/cargando.../i);
  expect(imageElement).toBeInTheDocument();
});
