# Action

Actions are a way to supercharge your alerts by gather relevant details about your system to better understand the alert-context.

These can be anything, from grabbing a quick chart to see overall CPU-utilization to getting a heap-dump from a crashing pod.

## How to build Actions

### Function Signature and Data-types

We build Actions by writing Action functions in golang which satisfy the type

```
func (types.Alert, types.Action, map[string]interface{})
```

This is the function that the tam-runtime will call once we register this Action function

The Action type looks like this

```
type Action struct {
	StepName       string `yaml:"step_name"`
	ActionName string `yaml:"Action_name"`
	ActionArgs string `yaml:"Action_args"`
}
```

### Program context

Once we are inside the function-context, we are free to do anything we want. <br>
We have the entire alert that was used to trigger this Action.

For action-functions, we also get a map containing the output of the enrichments configured of this alert-pipeline if any.

**NOTE:** Action Functions do not share context with each other. Thus you **CANNOT** use the output of one Action in another Action.

### Registering the Action-function

Once we have defined out function, we have to register it with the tam-runtime. This is basically updating a in-memory map which stores all the Actions available and a string identifier that is used to identify this function.

The function in [Action.go](./Action.go) looks like this

```
func LoadActions() {
	actMap := GetActionMap()
	actMap.Add("NOOP_ACTION", NoopAction)
	actMap.Add("SendToSlack", SendToSlackAction)
}
```

we will add new entry in this function which like this

```
enr.Add("ActionIdentifier", ActionFunctionName)
```

Here,
`ActionIdentifier` is the string value that will be used to refer to this Action in the tam-config; `ActionFunctionName` is the name of the function that we defined
