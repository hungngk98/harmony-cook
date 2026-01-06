import Menu from "../shared/Menu";
import "./SuggestChords.css";
import { useEffect, useState } from 'react';

const fetchSuggestedChords = (notes) => fetch(
    "http://localhost:5000/api/suggestchords",
    {
        headers: { "Content-Type": "application/json" },
        method: 'POST',
        body: JSON.stringify({ notes })
    }
).then(res => res.json())

export default function SuggestChords() {
    const [notes, setNotes] = useState("");
    const [chordGroups, setChordGroups] = useState({});

    const handleSubmit = async (e) => {
        e.preventDefault();

        try {
            const apiRes = await fetchSuggestedChords(notes.split(' ').map(n => n.trim()).filter(n => n.length > 0))

            if (apiRes.erCode == 0) {
                setChordGroups(apiRes.data);
                if (sessionStorage) {
                    sessionStorage.setItem("lastSuggestedChords", JSON.stringify({ notes: notes, chordGroups: apiRes.data }))
                }
            } else if (apiRes.erCode == 1) {
                if (apiRes.message) alert(apiRes.message);
                else alert("server error")
            }
        } catch (error) {
            alert("unknown error");
        }
    }

    useEffect(() => {
        if (sessionStorage) {
            const lastSuggestedChordsJSON = sessionStorage.getItem("lastSuggestedChords");
            if (lastSuggestedChordsJSON) {
                const lastSuggestedChords = JSON.parse(lastSuggestedChordsJSON);
                if (lastSuggestedChords.notes && lastSuggestedChords.chordGroups) {
                    setNotes(lastSuggestedChords.notes);
                    setChordGroups(lastSuggestedChords.chordGroups);
                }
            }
        }
    }, []);

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
            <table id='chord-suggestions'>
                <thead>
                    <tr>
                        <td id='chord-suggestions__col--note-count'>Số note trùng</td>
                        <td id='chord-suggestions__col--chords'>Hợp âm gợi ý</td>
                    </tr>
                </thead>
                <tbody>
                    {Object.keys(chordGroups).sort((a, b) => b - a).map((noteCount, i) => (
                        <tr key={i}>
                            <td>{noteCount}</td>
                            <td>{chordGroups[noteCount].join(', ')}</td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </div>
    )
}