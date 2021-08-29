# docker-landing-page

Simple landing page with links to running containers.
Useful as overview for local docker system (e.g. Synology NAS).

## Usage

1. Start landing page:
   ```
   docker run -d -v /var/run/docker.sock:/var/run/docker.sock -p8080:8080 ghcr.io/bboehmke/docker-landing-page
   ```
2. Add labels to container:
   * CLI:
   ```
   [...] --label landing-page.enabled=true --label landing-page.port=8000 --label landing-page.name=custom_name [...]
   ```
   * Docker compose:
   ```yaml
   labels:
      landing-page.enabled: true # Enables landing page for this container 
      landing-page.port: 8000 # port of application
      landing-page.name: "custom_name" # Name of link
   ```
3. Open browser at http://127.0.0.1:8080

