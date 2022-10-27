<script lang="js">
  import Button, { Label } from '@smui/button';
  import Textfield from "@smui/textfield";
  import Icon from "@smui/textfield/icon";
  import HelperText from "@smui/textfield/helper-text";
  import IndexTabs from "./basicComponents/indexTabs.svelte"

  let focused = false;
  let value = null;
  let dirty = false;
  let invalid = false;
  $: disabled = focused || !value || !dirty || invalid;

  function clickHandler() {
    alert(`Sending to ${value}!`);
    value = null;
    dirty = false;
  }

  //initial count value passed in from Go
  export let InitCounterVal;

  let count = InitCounterVal;
  const increment = () => {
    count += 1;
  }

  
</script>


<h2>Example svelte material UI elements</h2>

<div class="margins">
  <h3>Email validator example</h3>
  <!--
      Note: when you bind to `invalid`, but you only want to
      monitor it instead of updating it yourself, you also
      should include `updateInvalid`.
    -->
  <Textfield
    type="email"
    bind:dirty
    bind:invalid
    updateInvalid
    bind:value
    label="To Address"
    style="min-width: 250px;"
    input$autocomplete="email"
    on:focus={() => (focused = true)}
    on:blur={() => (focused = false)}
    withTrailingIcon={!disabled}
  >
    <!--
        Since this icon is conditional, it needs to be wrapped
        in a fragment, and we need to provide withTrailingIcon.
      -->
    <svelte:fragment slot="trailingIcon">
      {#if !disabled}
        <Icon class="material-icons" role="button" on:click={clickHandler}
          >send</Icon
        >
      {/if}
    </svelte:fragment>
    <HelperText validationMsg slot="helper">
      That's not a valid email address.
    </HelperText>
  </Textfield>
</div>

<pre
  class="status">Focused: {focused}, Dirty: {dirty}, Invalid: {invalid}, Value: {value}</pre>

<hr>
<br/>
<br/>

<h3>Tabs example</h3>
<br/>
<IndexTabs></IndexTabs>

<hr>
<br/>
<br/>


<h3>Initial numerical value passed in from Go</h3>
<Button on:click={increment} variant="raised">
  Increment value: {count}
</Button>


<hr>
<br/>
<br/>

 



<style>
  .demo-cell {
    height: 60px;
    display: flex;
    justify-content: center;
    align-items: center;
    background-color: var(--mdc-theme-secondary, #333);
    color: var(--mdc-theme-on-secondary, #fff);
  }
</style>