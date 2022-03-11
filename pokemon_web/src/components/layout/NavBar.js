// NavBar is a component that returns a navbar for pokedex app.
const NavBar = () => {
  return (
    <nav className="navbar navbar-expand-lg navbar-light pokebar p-3 sticky grid grid-cols-4">
      <h1 className="text-2xl">Pokédex</h1>
      <div className="text-center"></div>
      <h1 className="text-2xl text-white col-span-2 text-center">Número</h1>
    </nav>
  );
};

export default NavBar;
