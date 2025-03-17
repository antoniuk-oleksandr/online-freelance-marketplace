/** @type {import('tailwindcss').Config} */
const config = {
    content: [
        "./index.html",
        "./src/**/*.svelte",
    ],
    safelist: [
        "bg-red-500",
        "bg-green-500",
        "bg-orange-500",
        "lg:min-w-header-dropdown-menu",
        "bg-gray-200",
        "text-gray-700",
        "border-gray-300",
        "bg-yellow-100",
        "text-yellow-700",
        "border-yellow-300",
        "bg-green-100",
        "text-green-700",
        "border-green-300",
        "bg-red-200",
        "text-red-800",
        "border-red-400",
        "bg-orange-100",
        "text-orange-700",
        "border-orange-300",
        "bg-red-100",
        "text-red-700",
        "border-red-300",
        "bg-blue-100",
        "text-blue-700",
        "border-blue-300",
        { pattern: /^size-\d+$/, variants: ['lg'] },
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
                    "0%": { opacity: "0" },
                    "100%": { opacity: "1" }
                },
                "fade-out": {
                    "0%": { opacity: "1" },
                    "100%": { opacity: "0" },
                },
                "drop-down": {
                    "0%": { opacity: "0", transform: "translateY(5%)" },
                    "100%": { opacity: "1", transform: "translateY(0%)" },
                },
                "drop-up": {
                    "0%": { opacity: "1", transform: "translateY(0%)" },
                    "100%": { opacity: "0", transform: "translateY(5%)" },
                },
                "toast": {
                    "0%": {
                        opacity: 0,
                        transform: "translateX(-20%)",
                        scale: "0.95"
                    },
                    "50%": {
                        opacity: 1,
                        transform: "translateX(0)",
                        scale: "1"
                    },
                    "100%": {
                        opacity: 1,
                        transform: "translateX(0)",
                        scale: "1"
                    },
                },
                "drawer-open": {
                    "0%": {
                        opacity: 0.7,
                        transform: "translateX(30%)",
                    },
                    "100%": {
                        opacity: 1,
                        transform: "translateX(0)",
                    },
                },
                "drawer-close": {
                    "0%": {
                        opacity: 0.7,
                        transform: "translateX(0)",
                    },
                    "100%": {
                        opacity: 1,
                        transform: "translateX(30%)",
                    },
                },
                toastExit: {
                    '0%': {
                        opacity: '1',
                        transform: 'translateX(0) scale(1)',
                    },
                    '50%': {
                        opacity: '1',
                        transform: 'translateX(0) scale(1)',
                    },
                    '100%': {
                        opacity: '0',
                        transform: 'translateX(-20%) scale(0.95)',
                    },
                },
            },
            animation: {
                "fade-in": "fade-in 700ms ease-out",
                "fade-out": "fade-out 700ms ease-out",
                "toast": "toast 500ms ease-out",
                "toastExit": "toastExit 500ms ease-out",
                "drawer-open": "drawer-open 350ms ease-out",
                "drawer-close": "drawer-close 350ms ease-out",
                "drop-down": "drop-down 350ms ease-out",
                "drop-up": "drop-up 350ms ease-out",
            },
            gridTemplateColumns: {
                service: "1fr 25rem",
                "sign-form": "36rem 1fr",
                "search-page": "1fr 23rem",
                "search-top-bar": "1fr auto",
                "orders-table": "5rem 6rem 1fr 13rem 10rem 10rem",
                "orders-table-mobile": "6rem 1fr"

            },
            spacing: {
                4.5: "1.125rem",
                packages: "calc(32px)",
                22: "5.5rem",
                10: "2.5rem",
            },
            width: {
                100: "25rem",
                112: "28rem",
                "search-sidebar": "23rem",
                48: "12rem",
            },
            minWidth: {
                "header-dropdown-menu": "15rem",
            },
            height: {
                form: "56rem",
                128: "32rem",
                "order-overview": "23.3125rem",
            },
            translate: {
                "search-sidebar": "23rem",
                112: "28rem",
            },
            maxWidth: {
                320: "80rem",
                256: "64rem",
                168: "42rem",
            },
            minHeight: {
                app: "calc(100vh - 4rem - 3rem - 4rem)",
                212: "53rem",
                "search-left-side-mobile": "calc(100vh - 4rem - 3rem)",
            },
            maxHeight: {
                "search-sidebar": "calc(100vh - 4rem - 3rem)",
                app: "calc(100vh - 12.1625rem)",
                "search-array": "calc(100vh - 4rem  - 3rem - 4rem - 6rem  - 3rem)",
                "search-array-mobile": "calc(100vh - 4rem - 6rem - 3rem)",
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
