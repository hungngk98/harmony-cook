import { Link } from "react-router";

export default function Menu() {
    return (
        <ul>
            <li><Link to="/suggesttones">Gợi ý tone</Link></li>
            <li><Link to="/suggestchords">Gợi ý hợp âm</Link></li>
        </ul>
    )
}