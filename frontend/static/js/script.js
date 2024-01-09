// frontend/static/js/script.js

document.getElementById('uploadForm').addEventListener('submit', async function (event) {
    event.preventDefault();

    const formData = new FormData(this);

    try {
        const response = await fetch('/api/upload', {
            method: 'POST',
            body: formData,
            headers: {
                'Accept': 'application/json',
            },
        });

        const data = await response.json();

        // Display the result
        document.getElementById('result').innerText = `Image uploaded successfully! URL: ${data.url}`;
    } catch (error) {
        console.error('Error uploading image:', error);
        document.getElementById('result').innerText = 'Error uploading image. Please try again.';
    }
});
