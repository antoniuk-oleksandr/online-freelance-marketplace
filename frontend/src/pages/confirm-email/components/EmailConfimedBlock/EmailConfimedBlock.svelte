<script lang="ts">
    import Icon from "@iconify/svelte";
    import EmailConfirmedBlockLayout
        from "@/pages/confirm-email/components/EmailConfimedBlock/EmailConfirmedBlockLayout.svelte";
    import Button from "@/common-components/Button/Button.svelte";
    import {Link} from "svelte-routing";
</script>

<EmailConfirmedBlockLayout>
    <Icon icon="material-symbols:check-circle" class="text-green-500" width="64" height="64"/>
    <h1 class="text-2xl font-bold">Email Confirmed</h1>
    <p class="text-center text-light-palette-text-secondary dark:text-dark-palette-text-secondary">
        Your email has been successfully confirmed.
        You can now proceed to use your account.
    </p>
    <p class="text-center text-light-palette-text-secondary dark:text-dark-palette-text-secondary">
        Thank you for verifying your email. We’re excited to have you on board!
    </p>
    <Link to="/">
        <Button
                styles="w-48"
        >Go Home
        </Button>
    </Link>
</EmailConfirmedBlockLayout>
