<script lang="ts">
    import Icon from "@iconify/svelte";

    type GoogleButtonProps = {
        googleButtonText: string;
    }

    const {googleButtonText}: GoogleButtonProps = $props();

    const onClick = () => {
        console.log("Google button clicked");
    }
</script>

<button
        onclick={() => onClick()}
        type={"button"}
        class="flex !h-12 font-semibold hover:bg-light-palette-action-hover dark:hover:bg-dark-palette-action-hover duration-200 ease-out active:scale-95 items-center justify-center border border-light-palette-divider dark:border-dark-palette-divider rounded-md  gap-x-2 bg-light-palette-background-block dark:bg-dark-palette-background-block w-full"
>
    <Icon icon="flat-color-icons:google" width="24" height="24" />
    <span>{googleButtonText}</span>
</button>
