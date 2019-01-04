#### `alert(title, content)`
Creates an shows an alert modal to the user. A preview of how it may look can be found [here](https://i.zackrauen.com/7qnnNC.png).

|Parameter|Type|Description|
|-|-|:-|
|`title`|string|The title to show on the modal.|
|`content`|string|Content to show in the modal (can be html string).|

***

#### `clearCSS(id)`
Removes a style added with [`injectCSS`](#injectcssid-css) below.

|Parameter|Type|Description|
|-|-|:-|
|`id`|string|ID of the node to remove.|

***

#### `deleteData(pluginName, key)`
Deletes some saved data for plugin `pluginName` with key `key`.

|Parameter|Type|Description|
|-|-|:-|
|`pluginName`|string|Which plugin this is being used for.|
|`key`|string|Key for which data should be deleted.|

***

#### `findModule(filter)`
Searches for an internal Discord webpack module based on `filter`.

|Parameter|Type|Description|
|-|-|:-|
|`filter`|function|A function to use to filter modules.|

**@Returns** `{any|null}` - the module found or null if none were found.

***

#### `findAllModules(filter)`
Searches for multiple internal Discord webpack module based on `filter`. It's the same as [`findModule`](#findmodulefilter) but will return all matches

|Parameter|Type|Description|
|-|-|:-|
|`filter`|function|A function to use to filter modules.|

**@Returns** `{Array<any>}` - the modules found or null if none were found.

***

#### `findModuleByProps(...props)`
Searches for an internal Discord webpack module that has every property passed.

|Parameter|Type|Description|
|-|-|:-|
|`props`|...string|A series of properties to check for.|

**@Returns** `{any|null}` - the module found or null if none were found.

***

#### `getCore()`
Returns BandagedBD's instance of the core module. Only use this if you know what you are doing.

**@Returns** `{Core}` - BBD's instantiated core module.

***

#### `getData(pluginName, key)`
Alias for [loadData(pluginName, key)](#loaddatapluginname-key)

***

#### `getInternalInstance(node)`
Gets the internal react instance for a particular node.

|Parameter|Type|Description|
|-|-|:-|
|`node`|HTMLElement|jQuery|Node to find the react instance for.|

**@Returns** `{object|undefined}` - the instance if found or undefined otherwise.

***

#### `getPlugin(name)`
Gets the instance of another plugin with the name `name`.

|Parameter|Type|Description|
|-|-|:-|
|`name`|string|Name of the plugin to retreive.|

**@Returns** `{object|null}` - the plugin if found or null otherwise.

***

#### `injectCSS(id, css)`
Adds a block of css to the current document's `head`.

|Parameter|Type|Description|
|-|-|:-|
|`id`|string|Identifier for the node to be added. Can be used later with [`clearCSS`](#clearcssid) from above.|
|`css`|string|String of css to be added.|

**@Returns** `{object|null}` - the plugin if found or null otherwise.

***

#### `linkJS(id, url)`
Links some remote JavaScript to be added to the page. Useful for libraries like `Sortable.js`.

|Parameter|Type|Description|
|-|-|:-|
|`id`|string|Identifier for the node to be added. Can be used later with [`unlinkJS`](#unlinkjsid) below.|
|`url`|string|URL of the js.|

***

#### `loadData(pluginName, key)`
Gets some saved data for plugin `pluginName` with key `key`. Data can be saved with [`saveData`](#savedatapluginname-key-data).

|Parameter|Type|Description|
|-|-|:-|
|`pluginName`|string|Which plugin this is being used for.|
|`key`|string|Key for which data should be returned.|

**@Returns** `{any|null}` - The information that was saved previously, or null otherwise.

***


#### `monkeyPatch(module, methodName, options)`
This function monkey-patches a method on an object. The patching callback may be run before, after or instead of target method.
 - Be careful when monkey-patching. Think not only about original functionality of target method and your changes, but also about developers of other plugins, who may also patch this method before or after you. Try to change target method behaviour as little as possible, and avoid changing method signatures.
 - Display name of patched method is changed, so you can see if a function has been patched (and how many times) while debugging or in the stack trace. Also, patched methods have property `__monkeyPatched` set to `true`, in case you want to check something programmatically.

|Parameter|Type|Description|
|-|-|:-|
|`module`|object|Object to be patched. You can can also pass class prototypes to patch all class instances.|
|`methodName`|string|The name of the target message to be patched.|
|`options`|object|Options object. You should provide at least one of `before`, `after` or `instead` parameters. Other parameters are optional.|
|`[options.once=false]`|boolean|Set to `true` if you want to automatically unpatch method after first call.|
|`[options.silent=false]`|boolean|Set to `true` if you want to suppress log messages about patching and unpatching. Useful to avoid clogging the console in case of frequent conditional patching/unpatching, for example from another monkeyPatch callback.|
|`[options.displayName]`|string|You can provide meaningful name for class/object provided in `what` param for logging purposes. By default, this function will try to determine name automatically.|
|`options.before`|[PatchFunction](#callback-patchfunction)|Callback that will be called before original target method call. You can modify arguments here, so it will be passed to original method. Can be combined with `after`.|
|`options.after`|[PatchFunction](#callback-patchfunction)|Callback that will be called after original target method call. You can modify return value here, so it will be passed to external code which calls target method. Can be combined with `before`.|
|`options.instead`|[PatchFunction](#callback-patchfunction)|Callback that will be called instead of original target method call. You can get access to original method using `originalMethod` parameter if you want to call it, but you do not have to. Can't be combined with `before` and `after`.|

**@Returns** [`{CancelPatch}`](#callback-patchfunction) - A cancel function which allows you to undo the patch.

***

#### `onRemoved(node, callback)`
Adds a listener for when the node is removed from the document body.

|Parameter|Type|Description|
|-|-|:-|
|`node`|HTMLElement|Node to wait for.|
|`callback`|function|Function to be performed on event.|

***

#### `saveData(pluginName, key, data)`
Saved some `data` for plugin `pluginName` under `key` key. Gets saved in the plugins folder under `pluginName.config.json`. Data can be saved with [`loadData`](#loaddatapluginname-key).

|Parameter|Type|Description|
|-|-|:-|
|`pluginName`|string|Which plugin this is being used for.|
|`key`|string|Key for the data should be saved under.|
|`data`|any|Data to save.|

***

#### `setData(pluginName, key, data)`
Alias for [saveData(pluginName, key, data)](#savedatapluginname-key-data)

***

#### `showToast(content, options = {})`
Shows a simple toast message similar to on Android. An example of the `success` toast can be seen [here](https://i.zackrauen.com/zIagVa.png).

|Parameter|Type|Description|
|-|-|:-|
|`content`|string|Content to show inside the toast.|
|`[options]`|object|Options for the toast.|
|`[options.type=""]`|string|Changes the type of the toast stylistically and semantically. Choices: "", "info", "success", "danger"/"error", "warning"/"warn". Default: ""|
|`[options.icon=true]`|boolean|Determines whether the icon should show corresponding to the type. A toast without type will always have no icon. Default: true|
|`[options.timeout=3000]`|number|Adjusts the time (in ms) the toast should be shown for before disappearing automatically. Default: 3000|

***

#### `suppressErrors(method, message)`
Wraps a function in a try catch block.

|Parameter|Type|Description|
|-|-|:-|
|`method`|function|Function to wrap.|
|`[message]`|string|Additional info for any errors.|

**@Returns** `{function}` - The wrapped version of the original function.

***

#### `testJSON(data)`
Determines if the input is valid and parseable JSON.

|Parameter|Type|Description|
|-|-|:-|
|`data`|string|Data to test.|

**@Returns** `{boolean}` - True if the data is valid, false otherwise.

***

#### `unlinkJS(id)`
Removes some previously linked JS by [`linkJS`](#linkjsid-url).

|Parameter|Type|Description|
|-|-|:-|
|`id`|string|ID of the node to remove.|
