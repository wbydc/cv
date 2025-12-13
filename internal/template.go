package internal

const HTMLTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>{{ .Name }} - CV</title>
  <style>
    :root {
      /* Standard Screen Sizes */
      --primary-color: #2c3e50;
      --accent-color: #3498db;
      --text-color: #333;
      --light-text: #666;
      --bg-color: #fff;
      --sidebar-bg: #f4f7f6;
      --border-color: #e0e0e0;
    
      /* Sizing Variables (Screen) */
      --base-font-size: 16px;
      --header-font-size: 2.5rem;
      --sidebar-padding: 30px 20px;
      --main-padding: 40px 30px;
      --gap-size: 20px;
      --job-margin: 25px;
    }

    * { box-sizing: border-box; margin: 0; padding: 0; }
  
    body {
      font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
      font-size: var(--base-font-size);
      color: var(--text-color);
      background: #555;
      display: flex;
      justify-content: center;
      padding: 20px;
    }

    .page {
      background: var(--bg-color);
      width: 210mm;
      min-height: 297mm;
      display: grid;
      grid-template-columns: 28% 72%;
      box-shadow: 0 0 10px rgba(0,0,0,0.3);
      position: relative;
    }

    /* Sidebar */
    .sidebar {
      background: var(--sidebar-bg);
      padding: var(--sidebar-padding);
      border-right: 1px solid var(--border-color);
      display: flex;
      flex-direction: column;
      gap: var(--gap-size);
      word-wrap: break-word; 
    }

    .photo-container {
      width: 120px;
      height: 120px;
      border-radius: 50%;
      overflow: hidden;
      margin: 0 auto 10px;
      border: 3px solid var(--accent-color);
    }

    .photo-container img {
      width: 100%;
      height: 100%;
      object-fit: cover;
    }

    .contact-section h4 {
      text-transform: uppercase;
      font-size: 0.85rem;
      color: var(--accent-color);
      margin-bottom: 8px;
      border-bottom: 2px solid var(--border-color);
      padding-bottom: 3px;
    }

    .contact-item {
      margin-bottom: 8px;
      font-size: 0.85rem;
      word-break: break-word; 
      hyphens: none;
      line-height: 1.3;
    }

    .contact-item .label {
      font-weight: bold;
      display: block;
      color: var(--primary-color);
      font-size: 0.8rem;
    }

    .contact-item a {
      color: var(--text-color);
      text-decoration: none;
    }

    /* Main Content */
    .main-content {
      padding: var(--main-padding);
    }

    .header {
      margin-bottom: 20px;
      border-bottom: 2px solid var(--primary-color);
      padding-bottom: 10px;
    }

    .name {
      font-size: var(--header-font-size);
      color: var(--primary-color);
      line-height: 1.1;
      white-space: nowrap; 
    }

    .title {
      font-size: 1.1rem;
      color: var(--accent-color);
      margin-top: 5px;
    }

    .section-title {
      font-size: 1.2rem;
      color: var(--primary-color);
      border-bottom: 1px solid var(--border-color);
      padding-bottom: 4px;
      margin-bottom: 15px;
      margin-top: 5px;
      break-after: avoid;
    }

    /* Job Cards */
    .job-card {
      margin-bottom: var(--job-margin);
      break-inside: avoid; 
      page-break-inside: avoid;
    }

    .job-header {
      margin-bottom: 4px;
    }

    .job-title-row {
      display: flex;
      justify-content: space-between;
      align-items: baseline;
    }

    .job-position {
      font-size: 1rem;
      font-weight: bold;
      color: var(--primary-color);
    }

    .job-date {
      font-size: 0.8rem;
      color: var(--light-text);
      font-weight: 500;
      white-space: nowrap;
    }

    .job-company-row {
      font-size: 0.85rem;
      color: var(--light-text);
    }

    .job-company {
      font-weight: 600;
      color: var(--text-color);
    }

    .job-description, .job-achievements {
      margin-bottom: 10px;
      font-size: 0.95rem;
      line-height: 1.5;
    }
  
    .job-achievements {
      margin-left: 20px;
    }

    .job-achievements li {
      margin-bottom: 4px;
      break-inside: avoid;
    }

    .job-skills {
      display: flex;
      flex-wrap: wrap;
      gap: 4px;
      break-inside: avoid;
    }

    .skill-tag {
      background: #eef2f5;
      color: var(--primary-color);
      padding: 1px 6px;
      border-radius: 3px;
      font-size: 0.75rem;
      font-weight: 500;
      white-space: nowrap;
    }

    /* Floating Download Button */
    .download-btn-container {
      position: fixed;
      bottom: 30px;
      right: 30px;
      z-index: 9999;
    }

    .download-btn {
      background-color: var(--accent-color);
      color: white;
      border: none;
      padding: 15px 25px;
      border-radius: 50px;
      font-size: 16px;
      font-weight: bold;
      cursor: pointer;
      box-shadow: 0 4px 15px rgba(0,0,0,0.3);
      display: flex;
      align-items: center;
      gap: 10px;
    }

    .download-btn svg { width: 20px; height: 20px; fill: currentColor; }

    /* --- PRINT OPTIMIZATION --- */
    @page {
      margin: 0;
      size: auto;
    }

    @media print {
      :root {
        --base-font-size: 12px;
        --header-font-size: 1.8rem;
        --sidebar-padding: 20px 15px;
        --main-padding: 25px 20px;
        --gap-size: 15px;
        --job-margin: 15px;
      }

      .no-print { display: none !important; }

      html, body {
        height: 100%;
        width: 100%;
        margin: 0 !important;
        padding: 0 !important;
        background: none;
        display: block;
        -webkit-print-color-adjust: exact;
        print-color-adjust: exact;
      }

      .page {
        width: 100%;
        height: 100%;
        margin: 0 !important;
        box-shadow: none;
        grid-template-columns: 28% 72%;
      }

      .sidebar {
        background-color: var(--sidebar-bg) !important;
        -webkit-print-color-adjust: exact;
        height: 100%;
      }

      .skill-tag {
        background-color: #eef2f5 !important;
        -webkit-print-color-adjust: exact;
        border: 1px solid #ddd;
      }
    
      .photo-container {
        width: 100px;
        height: 100px;
      }

      .job-description { line-height: 1.25; }
      .contact-item { line-height: 1.2; }
    
      * { hyphens: none !important; }
    }
  </style>
</head>
<body>
  <!-- Floating Download Button -->
  <div class="download-btn-container no-print">
    <button class="download-btn" onclick="window.print()">
      <svg viewBox="0 0 24 24">
        <path d="M19 9h-4V3H9v6H5l7 7 7-7zM5 18v2h14v-2H5z"/>
      </svg>
      Save as PDF
    </button>
  </div>

  <div class="page">
    <!-- Sidebar -->
    <aside class="sidebar">
      {{ if .PhotoBase64 }}
      <div class="photo-container">
        <img src="{{ .PhotoBase64 }}" alt="{{ .Name }}">
      </div>
      {{ end }}

      <div class="contact-section">
        <h4>Contact</h4>
       
        <div class="contact-item">
          <span class="label">Email</span>
          <span class="protected-data" data-type="email" data-value="{{ .EmailEncoded }}">
            [Protected]
          </span>
        </div>

        <div class="contact-item">
          <span class="label">Phone</span>
          <span class="protected-data" data-type="phone" data-value="{{ .PhoneEncoded }}">
            [Protected]
          </span>
        </div>

        <div class="contact-item">
          <span class="label">Location</span>
          {{ .Location }}
        </div>
      </div>

      <div class="contact-section">
        <h4>Socials</h4>
        {{ range $platform, $url := .Socials }}
        <div class="contact-item">
          <span class="label">{{ $platform }}:</span>
          <a href="{{ $url }}" target="_blank">{{ $url }}</a>
        </div>
        {{ end }}
      </div>

      <div class="contact-section">
        <h4>Languages</h4>
        {{ range $lang, $level := .Languages }}
        <div class="contact-item">
          <span class="label">{{ $lang }}</span>
          {{ $level }}
        </div>
        {{ end }}
      </div>
    </aside>

    <!-- Main Content -->
    <main class="main-content">
      <header class="header">
        <h1 class="name">{{ .Name }}</h1>
        {{ if gt (len .Experience) 0 }}
        <div class="title">{{ (index .Experience 0).Position }}</div>
        {{ end }}
      </header>

      <section>
        <h2 class="section-title">Experience</h2>
        {{ range .Experience }}
        <div class="job-card">
          <div class="job-header">
            <div class="job-title-row">
              <h3 class="job-position">{{ .Position }}</h3>
              <span class="job-date">{{ .StartDate }} – {{ .EndDate }}</span>
            </div>
            <div class="job-company-row">
              <span class="job-company">{{ .Company }}</span>
              <span class="job-details">
                {{ .Location }} {{ if .IsRemote }}(Remote){{ end }} • {{ .Type }}
              </span>
            </div>
          </div>
          <div class="job-description">
            {{ range .Description }}<p>{{ . }}</p>{{ end }}
          </div>
          <ul class="job-achievements">
            {{ range .Achievements }}<li>{{ . }}</li>{{ end }}
          </ul>
          <div class="job-skills">
            {{ range .Skills }}<span class="skill-tag">{{ . }}</span>{{ end }}
          </div>
        </div>
        {{ end }}
      </section>
    </main>
  </div>

  <script>
    document.addEventListener('DOMContentLoaded', () => {
      const protectedElements = document.querySelectorAll('.protected-data');
      
      protectedElements.forEach(el => {
        const type = el.getAttribute('data-type');
        const encoded = el.getAttribute('data-value');
        
        if (encoded) {
          try {
            const decoded = atob(encoded);
            if (type === 'email') {
              el.innerHTML = '<a href="mailto:' + decoded + '">' + decoded + '</a>';
            } else if (type === 'phone') {
              el.innerHTML = '<a href="tel:' + decoded + '">' + decoded + '</a>';
            }
          } catch (e) {
            console.error('Failed to decode contact info');
          }
        }
      });
    });
  </script>
</body>
</html>
`
