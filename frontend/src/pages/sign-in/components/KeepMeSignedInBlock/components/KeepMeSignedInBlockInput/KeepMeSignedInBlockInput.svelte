<script lang="ts">
    type KeepMeSignedInBlockInputProps = {
        keepSignedIn: boolean,
    }

    const {keepSignedIn}: KeepMeSignedInBlockInputProps = $props();
</script>

<input
        class="hidden"
        checked={keepSignedIn}
        type="checkbox"
        name="keepSignedIn"
/>