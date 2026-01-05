import Menu from "../shared/Menu";
import "./SuggestTones.css";
import { useState } from "react";

const fetchSuggestedTones = (notes) => fetch(
    "http://localhost:5000/api/suggesttones",
    {
        headers: { "Content-Type": "application/json" },
        method: 'POST',
        body: JSON.stringify({ notes })
    }
).then(res => res.json())

export default function SuggestTones() {
    const [notes, setNotes] = useState("");
    const [toneGroups, setToneGroups] = useState([]);

    const handleSubmit = async (e) => {
        e.preventDefault();

        try {
            const apiRes = await fetchSuggestedTones(notes.split(' ').map(n => n.trim()).filter(n => n.length > 0))

            if (apiRes.erCode == 0) {
                setToneGroups(apiRes.data);
            } else if (apiRes.erCode == 1) {
                if (apiRes.message) alert(apiRes.message);
                else alert("server error")
            }
        } catch (error) {
            alert("unknown error");
        }
    }

    return (
        <div>
            <Menu />
            <form onSubmit={handleSubmit}>
                <div>
                    <label htmlFor="input--notes">Notes: </label>
                    <input id='input--notes' type="text" value={notes} onChange={(e) => setNotes(e.target.value)} placeholder='D Gb C# ...' />
                </div>
                <button type='submit'>Tìm</button>
            </form>
            <table id='tone-suggestions'>
                <thead>
                    <tr>
                        <td id='tone-suggestions__col--note-count'>Số note trùng</td>
                        <td id='tone-suggestions__col--tones'>Tone gợi ý</td>
                    </tr>
                </thead>
                <tbody>
                    {Object.keys(toneGroups).sort((a, b) => b - a).map((noteCount, i) => (
                        <tr key={i}>
                            <td>{noteCount}</td>
                            <td>{toneGroups[noteCount].join(', ')}</td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </div>
    )
}