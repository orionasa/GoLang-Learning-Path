import { useState, useEffect } from 'react'
import { Routes, Route, Link, useParams } from 'react-router-dom'
import ReactMarkdown from 'react-markdown'

function LessonsList() {
  const [lessons, setLessons] = useState([])

  useEffect(() => {
    fetch('/api/lessons')
      .then(res => res.json())
      .then(data => setLessons(data))
  }, [])

  return (
    <div>
      <h1>Go Lessons</h1>
      <ul>
        {lessons.map(lesson => (
          <li key={lesson}>
            <Link to={`/lessons/${lesson}`}>{lesson}</Link>
          </li>
        ))}
      </ul>
    </div>
  )
}

function Lesson() {
  const { filename } = useParams()
  const [content, setContent] = useState('')

  useEffect(() => {
    fetch(`/api/lessons/${filename}`)
      .then(res => res.text())
      .then(data => setContent(data))
  }, [filename])

  return (
    <div>
      <ReactMarkdown>{content}</ReactMarkdown>
    </div>
  )
}

function App() {
  return (
    <Routes>
      <Route path="/" element={<LessonsList />} />
      <Route path="/lessons/:filename" element={<Lesson />} />
    </Routes>
  )
}

export default App
