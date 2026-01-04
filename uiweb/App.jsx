import { BrowserRouter, Routes, Route } from "react-router";
import './App.css';
import SuggestChords from "./modules/SuggestChords/SuggestChords";
import SuggestScales from "./modules/SuggestScales/SuggestScales";
import Menu from "./modules/shared/Menu";

export default function App() {
    return (
        <BrowserRouter>
            <Routes>
                <Route path="/" Component={Home} />
                <Route path="/suggestchords" Component={SuggestChords} />
                <Route path="/suggestscales" Component={SuggestScales} />
            </Routes>
        </BrowserRouter>
    )
}

function Home() {
    return (
        <div>
            <Menu />
        </div>
    )
}