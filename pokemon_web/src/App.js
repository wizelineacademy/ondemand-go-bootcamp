import "./App.css";
import NavBar from "./components/layout/NavBar";
import PokemonSelector from "./components/PokemonSelector";
import PokemonDetail from "./components/PokemonDetail";
import { useEffect, useState, useRef } from "react";
import { GetPokemons } from "./api/PokemonService";

// App is a component that returns a simple pokedex app.
function App() {
  const [pokemons, setPokemons] = useState([]);
  const [selectedPokemon, setSelectedPokemon] = useState(null);

  const detailsRef = useRef(null);
  useEffect(() => {
    GetPokemons().then((data) => {
      setPokemons(data);
    });
  }, []);

  useEffect(() => {
    document.addEventListener("keydown", handleKeyDown);
    return () => {
      document.removeEventListener("keydown", handleKeyDown);
    };
  }, [pokemons, selectedPokemon]);

  const handleKeyDown = (e) => {
    switch (e.code) {
      case "ArrowUp":
        handleMoveSelection("up");
        break;
      case "ArrowDown":
        handleMoveSelection("down");
        break;
      default:
        break;
    }
  };

  const handleMoveSelection = (direction) => {
    const index = selectedPokemon
      ? pokemons.findIndex((pokemon) => pokemon.Id === selectedPokemon.Id)
      : 0;
    if (direction === "up" && index > 0) {
      setSelectedPokemon(pokemons[index - 1]);
    } else if (direction === "down" && index < pokemons.length - 1) {
      setSelectedPokemon(pokemons[index + 1]);
    } else {
      setSelectedPokemon(pokemons[0]);
    }
    detailsRef.current
      .querySelector(".pokemon-selected")
      .scrollIntoViewIfNeeded(false);
  };

  const handlePokemonSelected = (pokemon) => {
    setSelectedPokemon(pokemon);
  };

  return (
    <div className="App">
      <NavBar />
      <div className="w-full grid grid-cols-5 h-[calc(100%-4rem)]">
        <div className="details col-span-3">
          {selectedPokemon && <PokemonDetail image={selectedPokemon.Url} />}
        </div>
        <div
          ref={detailsRef}
          className="pokemon-selector col-span-2 max-h-screen overflow-x-auto"
        >
          {pokemons.length > 0 &&
            pokemons.map((p, idx) => (
              <PokemonSelector
                key={idx}
                number={p.Id}
                name={p.Name}
                image={p.Sprite}
                selected={p.Id === selectedPokemon?.Id}
                onClick={() => handlePokemonSelected(p)}
              />
            ))}
        </div>
      </div>
    </div>
  );
}

export default App;
