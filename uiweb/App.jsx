import { useState } from 'react';
import './App.css';

const fetchChordsRanked = (notes) => fetch(
    "http://localhost:5000/api/suggestchords",
    {
        headers: { "Content-Type": "application/json" },
        method: 'POST',
        body: JSON.stringify({ notes })
    }
).then(res => res.json())

export default function App() {
    const [notes, setNotes] = useState("");
    const [chordGroups, setChordGroups] = useState([]);

    const handleSubmit = async (e) => {
        e.preventDefault();

        const apiRes = await fetchChordsRanked(notes.split(',').map(n => n.trim()))

        if (apiRes.erCode == 0) {
            setChordGroups(apiRes.data);
        } else {
            alert("error");
        }
    }

    return (
        <div>
            <form onSubmit={handleSubmit}>
                <div>
                    <label htmlFor="input--notes">Notes: </label>
                    <input id='input--notes' type="text" value={notes} onChange={(e) => setNotes(e.target.value)} />
                </div>
                <button type='submit'>Search</button>
            </form>
            <table id='chord-suggestions'>
                <thead>
                    <tr>
                        <td id='chord-suggestions__col--note-count'>Note count</td>
                        <td id='chord-suggestions__col--chords'>Chords</td>
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