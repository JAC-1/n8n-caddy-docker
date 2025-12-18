{{$nothing := "There is nothing here"}}
<!DOCTYPE html>
<html lang="en">

{{include "includes/head.html" }}
{{include "includes/header.html" }}


<main class="max-w-7xl mx-auto px-4 py-8">
  <h1>It works?</h1>
  {{include "includes/upload-form.html" }}
  {{include "includes/chat.html" }}

</main>

<script>
  window.__unocss = {
    rules: [
      // custom rules...
    ],
    presets: [
      // custom presets...
    ],
    // ...
  }
</script>

</html>
