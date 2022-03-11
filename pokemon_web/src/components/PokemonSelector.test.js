import { render, screen } from "@testing-library/react";
import PokemonSelector from "./PokemonSelector";

describe("Testing states", () => {
  test("render unselected pokemon", () => {
    const { container } = render(
      <PokemonSelector
        number={1}
        name="Bulbasaur"
        image="https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/1.png"
        selected={false}
        onClick={() => null}
      />
    );

    const mustNotExists = container.getElementsByClassName("pokemon-selected");

    expect(mustNotExists.length).toBe(0);
  });
  test("render selected pokemon", () => {
    const { container } = render(
      <PokemonSelector
        number={1}
        name="Bulbasaur"
        image="https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/1.png"
        selected={true}
        onClick={() => null}
      />
    );

    const mustExists = container.getElementsByClassName("pokemon-selected");
    expect(mustExists.length).toBe(1);
  });
});
