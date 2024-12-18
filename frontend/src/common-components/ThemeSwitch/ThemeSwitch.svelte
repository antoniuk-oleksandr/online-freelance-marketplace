<script lang="ts">
    import Icon from '@iconify/svelte';
    import ThemeSwitchLayout from "@/common-components/ThemeSwitch/ThemeSwitchLayout.svelte";
    import {isDarkMode} from "@/common-stores/theme-storage.ts";

    let darkMode = $state<boolean | null>(null);

    isDarkMode.subscribe((value) => {
        darkMode = value;
    })

    const toggleTheme = () => {
        isDarkMode.set(!darkMode);
    };
</script>

{#if darkMode !== null}
    <ThemeSwitchLayout>
        <button
                type="button"
                class={"text-lg flex items-center cursor-pointer duration-200 ease-out"}
                onclick={toggleTheme}
        >
            <Icon class="size-8 text-light-palette-text-primary dark:text-dark-palette-text-primary"
                  icon={`hugeicons:${darkMode ? 'sun-03' : 'moon-02'}`}/>
        </button>
    </ThemeSwitchLayout>
{/if}
