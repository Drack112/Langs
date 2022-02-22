var resourcesSuffix = "";
var darkThemes = ["dark", "ayu"];
window.currentTheme = document.getElementById("themeStyle");
window.mainTheme = document.getElementById("mainThemeStyle");
var settingsDataset = (function () {
  var settingsElement = document.getElementById("default-settings");
  if (settingsElement === null) {
    return null;
  }
  var dataset = settingsElement.dataset;
  if (dataset === undefined) {
    return null;
  }
  return dataset;
})();
function getSettingValue(settingName) {
  var current = getCurrentValue("rustdoc-" + settingName);
  if (current !== null) {
    return current;
  }
  if (settingsDataset !== null) {
    var def = settingsDataset[settingName.replace(/-/g, "_")];
    if (def !== undefined) {
      return def;
    }
  }
  return null;
}
var localStoredTheme = getSettingValue("theme");
var savedHref = [];
function hasClass(elem, className) {
  return elem && elem.classList && elem.classList.contains(className);
}
function addClass(elem, className) {
  if (!elem || !elem.classList) {
    return;
  }
  elem.classList.add(className);
}
function removeClass(elem, className) {
  if (!elem || !elem.classList) {
    return;
  }
  elem.classList.remove(className);
}
function onEach(arr, func, reversed) {
  if (arr && arr.length > 0 && func) {
    var length = arr.length;
    var i;
    if (reversed) {
      for (i = length - 1; i >= 0; --i) {
        if (func(arr[i])) {
          return true;
        }
      }
    } else {
      for (i = 0; i < length; ++i) {
        if (func(arr[i])) {
          return true;
        }
      }
    }
  }
  return false;
}
function onEachLazy(lazyArray, func, reversed) {
  return onEach(Array.prototype.slice.call(lazyArray), func, reversed);
}
function hasOwnPropertyRustdoc(obj, property) {
  return Object.prototype.hasOwnProperty.call(obj, property);
}
function updateLocalStorage(name, value) {
  try {
    window.localStorage.setItem(name, value);
  } catch (e) {}
}
function getCurrentValue(name) {
  try {
    return window.localStorage.getItem(name);
  } catch (e) {
    return null;
  }
}
function switchTheme(styleElem, mainStyleElem, newTheme, saveTheme) {
  var fullBasicCss = "rustdoc" + resourcesSuffix + ".css";
  var fullNewTheme = newTheme + resourcesSuffix + ".css";
  var newHref = mainStyleElem.href.replace(fullBasicCss, fullNewTheme);
  if (saveTheme) {
    updateLocalStorage("rustdoc-theme", newTheme);
  }
  if (styleElem.href === newHref) {
    return;
  }
  var found = false;
  if (savedHref.length === 0) {
    onEachLazy(document.getElementsByTagName("link"), function (el) {
      savedHref.push(el.href);
    });
  }
  onEach(savedHref, function (el) {
    if (el === newHref) {
      found = true;
      return true;
    }
  });
  if (found) {
    styleElem.href = newHref;
  }
}
function useSystemTheme(value) {
  if (value === undefined) {
    value = true;
  }
  updateLocalStorage("rustdoc-use-system-theme", value);
  var toggle = document.getElementById("use-system-theme");
  if (toggle && toggle instanceof HTMLInputElement) {
    toggle.checked = value;
  }
}
var updateSystemTheme = (function () {
  if (!window.matchMedia) {
    return function () {
      var cssTheme = getComputedStyle(
        document.documentElement
      ).getPropertyValue("content");
      switchTheme(
        window.currentTheme,
        window.mainTheme,
        JSON.parse(cssTheme) || "light",
        true
      );
    };
  }
  var mql = window.matchMedia("(prefers-color-scheme: dark)");
  function handlePreferenceChange(mql) {
    if (getSettingValue("use-system-theme") !== "false") {
      var lightTheme = getSettingValue("preferred-light-theme") || "light";
      var darkTheme = getSettingValue("preferred-dark-theme") || "dark";
      if (mql.matches) {
        switchTheme(window.currentTheme, window.mainTheme, darkTheme, true);
      } else {
        switchTheme(window.currentTheme, window.mainTheme, lightTheme, true);
      }
    }
  }
  mql.addListener(handlePreferenceChange);
  return function () {
    handlePreferenceChange(mql);
  };
})();
if (getSettingValue("use-system-theme") !== "false" && window.matchMedia) {
  if (
    getSettingValue("use-system-theme") === null &&
    getSettingValue("preferred-dark-theme") === null &&
    darkThemes.indexOf(localStoredTheme) >= 0
  ) {
    updateLocalStorage("rustdoc-preferred-dark-theme", localStoredTheme);
  }
  updateSystemTheme();
} else {
  switchTheme(
    window.currentTheme,
    window.mainTheme,
    getSettingValue("theme") || "light",
    false
  );
}
