import { useEffect, useState } from "react";
import PropTypes from "prop-types";

// PokemonDetail component is used to display the details of a pokemon.
const PokemonDetail = ({ image }) => {
  const [loaded, setLoaded] = useState(false);
  const handleLoadImage = () => {
    setLoaded(true);
  };
  useEffect(() => {
    setLoaded(false);
  }, [image]);
  return (
    <div className="pokemon-detail  h-full flex justify-center items-center pb-10">
      {!loaded && <div>cargando...</div>}

      <img
        className="h-1/3"
        src={image}
        alt="pokemon"
        onLoad={handleLoadImage}
        style={{ visibility: loaded ? "visible" : "hidden" }}
      />
    </div>
  );
};

PokemonDetail.propTypes = {
  image: PropTypes.string.isRequired,
};

export default PokemonDetail;
