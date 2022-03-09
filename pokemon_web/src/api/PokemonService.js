import Api from "./ApiService"

export const GetPokemons = async () => {
  return await Api.get("http://localhost:8090/Pokemons");
  
}