import { defineConfig } from "vitepress";

export default defineConfig({
  title: "Geocoder ID",
  description: "Offline reverse geocoding for Indonesia",

  head: [["link", { rel: "icon", href: "/favicon.ico" }]],

  themeConfig: {
    logo: "/logo.svg",

    nav: [
      { text: "Guide", link: "/guide/" },
      { text: "API", link: "/api/" },
      { text: "GitHub", link: "https://github.com/zsbahtiar/geocoder-id" },
    ],

    sidebar: {
      "/guide/": [
        {
          text: "Getting Started",
          items: [
            { text: "Introduction", link: "/guide/" },
            { text: "Installation", link: "/guide/installation" },
            { text: "Quick Start", link: "/guide/quick-start" },
          ],
        },
      ],
      "/api/": [
        {
          text: "API Reference",
          items: [
            { text: "Overview", link: "/api/" },
            { text: "CLI", link: "/api/cli" },
            { text: "Go Library", link: "/api/go" },
          ],
        },
      ],
    },

    socialLinks: [
      { icon: "github", link: "https://github.com/zsbahtiar/geocoder-id" },
    ],

    footer: {
      message: "Released under the MIT License.",
      copyright: "Copyright © 2026 zsbahtiar",
    },

    search: {
      provider: "local",
    },
  },
});
