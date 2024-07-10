const textarea = document.getElementById('request');

textarea.addEventListener('input', function() {
    this.style.height = 'auto';
    this.style.height = (this.scrollHeight) + 'px';
});

document.body.addEventListener('htmx:afterSettle', function(evt) {
    if (evt.target.id === 'messages') {
        evt.target.scrollTop = evt.target.scrollHeight;
    }
});