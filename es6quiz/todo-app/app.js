const apiUrl = 'http://localhost:8000/notes';

const noteForm = document.getElementById('noteForm');
const noteId = document.getElementById('noteId');
const noteTitle = document.getElementById('noteTitle');
const noteContent = document.getElementById('noteContent');
const notesTable = document.getElementById('notesTable');
const submitBtn = document.getElementById('submitBtn');
const cancelBtn = document.getElementById('cancelBtn');
let notesList = [];

async function getNotes() {
  const response = await fetch(apiUrl);
  const notes = await response.json();
  notesList = notes;
  renderNotes(notes);
}

function renderNotes(notes) {
  notesTable.innerHTML = '';

  notes.forEach((note) => {
    notesTable.innerHTML += `
      <tr>
        <td>${note.id}</td>
        <td>${note.title}</td>
        <td>${note.content}</td>
        <td>
          <button class="btn btn-warning btn-sm me-2" data-action="edit" data-id="${note.id}">Edit</button>
          <button class="btn btn-danger btn-sm" data-action="delete" data-id="${note.id}">Delete</button>
        </td>
      </tr>
    `;
  });
}

async function addNote(note) {
  await fetch(apiUrl, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(note),
  });

  await getNotes();
}

async function deleteNote(id) {
  await fetch(`${apiUrl}/${id}`, {
    method: 'DELETE',
  });

  await getNotes();
}

function editNote(note) {
  noteId.value = note.id;
  noteTitle.value = note.title;
  noteContent.value = note.content;
  submitBtn.textContent = 'Update Note';
  cancelBtn.classList.remove('d-none');
}

async function updateNote(id, updatedNote) {
  await fetch(`${apiUrl}/${id}`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(updatedNote),
  });

  await getNotes();
}

function resetForm() {
  noteForm.reset();
  noteId.value = '';
  submitBtn.textContent = 'Add Note';
  cancelBtn.classList.add('d-none');
}

noteForm.addEventListener('submit', async function (event) {
  event.preventDefault();

  const note = {
    title: noteTitle.value,
    content: noteContent.value,
  };

  if (noteId.value) {
    await updateNote(noteId.value, note);
  } else {
    await addNote(note);
  }

  resetForm();
});

cancelBtn.addEventListener('click', resetForm);

notesTable.addEventListener('click', async function (event) {
  const button = event.target;
  const id = button.dataset.id;

  if (button.dataset.action === 'delete') {
    await deleteNote(id);
  }

  if (button.dataset.action === 'edit') {
    const note = notesList.find((item) => String(item.id) === id);
    editNote(note);
  }
});

getNotes();
