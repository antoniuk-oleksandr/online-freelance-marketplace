<script>
    import {navigate} from "svelte-routing";
    import Button from "@/common-components/Button/Button.svelte";
</script>

<Button
        clickAction={() => navigate("/sign-in")}
        styles="text-lg !text-light-palette-text-primary dark:!text-dark-palette-text-primary !bg-transparent hover:!opacity-100 hover:!bg-transparent !border !border-transparent hover:!border-light-palette-divider hover:dark:!border-dark-palette-divider"
>Return to Sign In
</Button>