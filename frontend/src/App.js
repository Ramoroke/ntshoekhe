import {Route, BrowserRouter as Router, Routes} from "react-router-dom"
import Home from "./pages/Home";
import AddDrug from "./pages/AddDrug";
import EditDrug from "./pages/EditDrug";

function App() {
  return (
    <Router>
      <Routes>
        <Route exact path="/" element={<Home />} />
        <Route path="/create" element={<AddDrug />} />
        <Route path="/update" element={<EditDrug />} />
      </Routes>
    </Router>
  );
}

export default App;
