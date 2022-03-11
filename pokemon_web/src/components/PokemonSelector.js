import PropTypes from "prop-types";
import pokeball from "../assets/img/pokeball.png";
import pokeblack from "../assets/img/pokeball-black.png";

// PokemonSelector component is used to display the list of pokemons.
const PokemonSelector = ({ number, name, image, selected = true, onClick }) => {
  return (
    <div
      className={
        selected
          ? "grid grid-cols-8 align-center bg-white rounded-full m-4 h-18 shadow-xl"
          : "grid grid-cols-8 align-center rounded-full m-4 h-12"
      }
      style={selected ? { background: "#fe5d13" } : {}}
      onClick={() => onClick(number)}
    >
      <div className="h-12 flex items-center justify-center">
        <img src={image} alt={name} />
      </div>
      <div className="col-span-3 h-16 flex items-center">
        <h1 className="text-xl">{`N.º ${number}`}</h1>
      </div>
      <div
        className={
          selected
            ? "pokemon-selected col-span-3 text-white text-xl h-16  flex items-center pl-6"
            : "col-span-3 text-xl h-16  flex items-center pl-6"
        }
        style={
          selected
            ? {
                background:
                  "linear-gradient(290deg, #303030 90%, #fff 15%, #fe5d13 10%, #fe5d13 10%)",
              }
            : {}
        }
      >
        <p>{name}</p>
      </div>
      {selected && (
        <div className="h-16 flex items-center rounded-r-full" style={{ background: "#303030" }}>
          <img  style={{height: "36px"}} src={pokeball} alt="pokeball" />
        </div>
      )}
      {!selected && <div className="h-16 flex items-center rounded-r-full" >
          <img style={{height: "42px"}} src={pokeblack} alt="pokeball" />
        </div>}
    </div>
  );
};

PokemonSelector.propTypes = {
  number: PropTypes.number.isRequired,
  name: PropTypes.string.isRequired,
  image: PropTypes.string.isRequired,
  selected: PropTypes.bool.isRequired,
  onClick: PropTypes.func.isRequired,
};

export default PokemonSelector;
