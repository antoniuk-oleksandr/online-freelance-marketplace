/** @type {import('tailwindcss').Config} */
const config = {
    content: [
        "./index.html",
        "./src/**/*.svelte",
    ],
    theme: {
        extend: {
            transitionProperty: {
                translate: 'transform',
            },
            transitionDuration: {
                'translate': '200ms',
            },
            keyframes: {
                "fade-in": {
                    "0%": {opacity: "0"},
                    "100%": {opacity: "1"}
                }
            },
            animation: {
                "fade-in": "fade-in 700ms ease-out",
            },
            gridTemplateColumns: {
                service: "1fr 25rem",
                "sign-form": "36rem 1fr",
            },
            spacing: {
                4.5: "1.125rem",
                packages: "calc(32px)"
            },
            width: {
                100: "25rem",
            },
            height: {
                form : "56rem",
            },
            maxWidth: {
                320: "80rem",
                256: "64rem",
            },
            minHeight: {
                app: "calc(100svh - 12.1625rem)",
                212: "53rem",
            },
            colors: {
                dark: {
                    palette: {
                        action: {
                            active: "#fff",
                            hover: "rgba(255, 255, 255, 0.08)",
                            selected: "rgba(255, 255, 255, 0.16)",
                            disabled: "rgba(255, 255, 255, 0.32)",
                            disabledBackground: "rgba(255, 255, 255, 0.12)",
                        },
                        text: {
                            primary: "#fff",
                            secondary: "rgba(255, 255, 255, 0.7)",
                            disabled: "rgba(255, 255, 255, 0.5)",
                        },
                        background: {
                            default: "#1a1a1a",
                            block: "#222222",
                        },
                        divider: "rgba(255, 255, 255, 0.2)",
                    }
                },
                light: {
                    palette: {
                        action: {
                            active: "#000",
                            hover: "rgba(0, 0, 0, 0.08)",
                            selected: "rgba(0, 0, 0, 0.16)",
                            disabled: "rgba(0, 0, 0, 0.32)",
                            disabledBackground: "rgba(0, 0, 0, 0.12)",
                        },
                        text: {
                            primary: "#000",
                            secondary: "rgba(0, 0, 0, 0.7)",
                            disabled: "rgba(0, 0, 0, 0.5)",
                        },
                        background: {
                            default: "#f9f9f9",
                            block: "#fff"
                        },
                        divider: "rgba(0, 0, 0, 0.2)",
                    }
                }
            }
        },
    },
    plugins: [],
    darkMode: "class",
};

export default config;
