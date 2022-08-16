package assets

var CpScriptsJs = []byte(`! function(t, e) {
	void 0 === t && void 0 !== window && (t = window), "function" == typeof define && define.amd ? define(["jquery"], function(t) {
		return e(t)
	}) : "object" == typeof module && module.exports ? module.exports = e(require("jquery")) : e(t.jQuery)
}(this, function(t) {
	! function(k) {
		"use strict";
		var d = ["sanitize", "whiteList", "sanitizeFn"],
			s = ["background", "cite", "href", "itemtype", "longdesc", "poster", "src", "xlink:href"],
			t = {
				"*": ["class", "dir", "id", "lang", "role", "tabindex", "style", /^aria-[\w-]*$/i],
				a: ["target", "href", "title", "rel"],
				area: [],
				b: [],
				br: [],
				col: [],
				code: [],
				div: [],
				em: [],
				hr: [],
				h1: [],
				h2: [],
				h3: [],
				h4: [],
				h5: [],
				h6: [],
				i: [],
				img: ["src", "alt", "title", "width", "height"],
				li: [],
				ol: [],
				p: [],
				pre: [],
				s: [],
				small: [],
				span: [],
				sub: [],
				sup: [],
				strong: [],
				u: [],
				ul: []
			},
			r = /^(?:(?:https?|mailto|ftp|tel|file):|[^&:/?#]*(?:[/?#]|$))/gi,
			l = /^data:(?:image\/(?:bmp|gif|jpeg|jpg|png|tiff|webp)|video\/(?:mpeg|mp4|ogg|webm)|audio\/(?:mp3|oga|ogg|opus));base64,[a-z0-9+/]+=*$/i;

		function g(t, e) {
			var n = t.nodeName.toLowerCase();
			if (-1 !== k.inArray(n, e)) return -1 === k.inArray(n, s) || Boolean(t.nodeValue.match(r) || t.nodeValue.match(l));
			for (var i = k(e).filter(function(t, e) {
					return e instanceof RegExp
				}), o = 0, a = i.length; o < a; o++)
				if (n.match(i[o])) return 1
		}

		function z(t, e, n) {
			if (n && "function" == typeof n) return n(t);
			for (var i = Object.keys(e), o = 0, a = t.length; o < a; o++)
				for (var s = t[o].querySelectorAll("*"), r = 0, l = s.length; r < l; r++) {
					var c = s[r],
						d = c.nodeName.toLowerCase();
					if (-1 !== i.indexOf(d))
						for (var u = [].slice.call(c.attributes), h = [].concat(e["*"] || [], e[d] || []), p = 0, f = u.length; p < f; p++) {
							var m = u[p];
							g(m, h) || c.removeAttribute(m.nodeName)
						} else c.parentNode.removeChild(c)
				}
		}
		"classList" in document.createElement("_") || function(t) {
			if ("Element" in t) {
				var e = "classList",
					n = "prototype",
					i = t.Element[n],
					o = Object,
					a = function() {
						var n = k(this);
						return {
							add: function(t) {
								return t = Array.prototype.slice.call(arguments).join(" "), n.addClass(t)
							},
							remove: function(t) {
								return t = Array.prototype.slice.call(arguments).join(" "), n.removeClass(t)
							},
							toggle: function(t, e) {
								return n.toggleClass(t, e)
							},
							contains: function(t) {
								return n.hasClass(t)
							}
						}
					};
				if (o.defineProperty) {
					var s = {
						get: a,
						enumerable: !0,
						configurable: !0
					};
					try {
						o.defineProperty(i, e, s)
					} catch (t) {
						void 0 !== t.number && -2146823252 !== t.number || (s.enumerable = !1, o.defineProperty(i, e, s))
					}
				} else o[n].__defineGetter__ && i.__defineGetter__(e, a)
			}
		}(window);
		var e, c, n = document.createElement("_");
		if (n.classList.add("c1", "c2"), !n.classList.contains("c2")) {
			var i = DOMTokenList.prototype.add,
				o = DOMTokenList.prototype.remove;
			DOMTokenList.prototype.add = function() {
				Array.prototype.forEach.call(arguments, i.bind(this))
			}, DOMTokenList.prototype.remove = function() {
				Array.prototype.forEach.call(arguments, o.bind(this))
			}
		}
		if (n.classList.toggle("c3", !1), n.classList.contains("c3")) {
			var a = DOMTokenList.prototype.toggle;
			DOMTokenList.prototype.toggle = function(t, e) {
				return 1 in arguments && !this.contains(t) == !e ? e : a.call(this, t)
			}
		}

		function u(t) {
			if (null == this) throw new TypeError;
			var e = String(this);
			if (t && "[object RegExp]" == c.call(t)) throw new TypeError;
			var n = e.length,
				i = String(t),
				o = i.length,
				a = 1 < arguments.length ? arguments[1] : void 0,
				s = a ? Number(a) : 0;
			s != s && (s = 0);
			var r = Math.min(Math.max(s, 0), n);
			if (n < o + r) return !1;
			for (var l = -1; ++l < o;)
				if (e.charCodeAt(r + l) != i.charCodeAt(l)) return !1;
			return !0
		}

		function I(t) {
			var e, n = [],
				i = t.selectedOptions;
			if (t.multiple)
				for (var o = 0, a = i.length; o < a; o++) e = i[o], n.push(e.value || e.text);
			else n = t.value;
			return n
		}
		n = null, String.prototype.startsWith || (e = function() {
			try {
				var t = {},
					e = Object.defineProperty,
					n = e(t, t, t) && e
			} catch (t) {}
			return n
		}(), c = {}.toString, e ? e(String.prototype, "startsWith", {
			value: u,
			configurable: !0,
			writable: !0
		}) : String.prototype.startsWith = u), Object.keys || (Object.keys = function(t, e, n) {
			for (e in n = [], t) n.hasOwnProperty.call(t, e) && n.push(e);
			return n
		}), HTMLSelectElement && !HTMLSelectElement.prototype.hasOwnProperty("selectedOptions") && Object.defineProperty(HTMLSelectElement.prototype, "selectedOptions", {
			get: function() {
				return this.querySelectorAll(":checked")
			}
		});
		var h = {
			useDefault: !1,
			_set: k.valHooks.select.set
		};
		k.valHooks.select.set = function(t, e) {
			return e && !h.useDefault && k(t).data("selected", !0), h._set.apply(this, arguments)
		};
		var $ = null,
			p = function() {
				try {
					return new Event("change"), !0
				} catch (t) {
					return !1
				}
			}();

		function S(t, e, n, i) {
			for (var o = ["display", "subtext", "tokens"], a = !1, s = 0; s < o.length; s++) {
				var r = o[s],
					l = t[r];
				if (l && (l = l.toString(), "display" === r && (l = l.replace(/<[^>]+>/g, "")), i && (l = y(l)), l = l.toUpperCase(), a = "contains" === n ? 0 <= l.indexOf(e) : l.startsWith(e))) break
			}
			return a
		}

		function A(t) {
			return parseInt(t, 10) || 0
		}
		k.fn.triggerNative = function(t) {
			var e, n = this[0];
			n.dispatchEvent ? (p ? e = new Event(t, {
				bubbles: !0
			}) : (e = document.createEvent("Event")).initEvent(t, !0, !1), n.dispatchEvent(e)) : n.fireEvent ? ((e = document.createEventObject()).eventType = t, n.fireEvent("on" + t, e)) : this.trigger(t)
		};
		var f = {
				"À": "A",
				"Á": "A",
				"Â": "A",
				"Ã": "A",
				"Ä": "A",
				"Å": "A",
				"à": "a",
				"á": "a",
				"â": "a",
				"ã": "a",
				"ä": "a",
				"å": "a",
				"Ç": "C",
				"ç": "c",
				"Ð": "D",
				"ð": "d",
				"È": "E",
				"É": "E",
				"Ê": "E",
				"Ë": "E",
				"è": "e",
				"é": "e",
				"ê": "e",
				"ë": "e",
				"Ì": "I",
				"Í": "I",
				"Î": "I",
				"Ï": "I",
				"ì": "i",
				"í": "i",
				"î": "i",
				"ï": "i",
				"Ñ": "N",
				"ñ": "n",
				"Ò": "O",
				"Ó": "O",
				"Ô": "O",
				"Õ": "O",
				"Ö": "O",
				"Ø": "O",
				"ò": "o",
				"ó": "o",
				"ô": "o",
				"õ": "o",
				"ö": "o",
				"ø": "o",
				"Ù": "U",
				"Ú": "U",
				"Û": "U",
				"Ü": "U",
				"ù": "u",
				"ú": "u",
				"û": "u",
				"ü": "u",
				"Ý": "Y",
				"ý": "y",
				"ÿ": "y",
				"Æ": "Ae",
				"æ": "ae",
				"Þ": "Th",
				"þ": "th",
				"ß": "ss",
				"Ā": "A",
				"Ă": "A",
				"Ą": "A",
				"ā": "a",
				"ă": "a",
				"ą": "a",
				"Ć": "C",
				"Ĉ": "C",
				"Ċ": "C",
				"Č": "C",
				"ć": "c",
				"ĉ": "c",
				"ċ": "c",
				"č": "c",
				"Ď": "D",
				"Đ": "D",
				"ď": "d",
				"đ": "d",
				"Ē": "E",
				"Ĕ": "E",
				"Ė": "E",
				"Ę": "E",
				"Ě": "E",
				"ē": "e",
				"ĕ": "e",
				"ė": "e",
				"ę": "e",
				"ě": "e",
				"Ĝ": "G",
				"Ğ": "G",
				"Ġ": "G",
				"Ģ": "G",
				"ĝ": "g",
				"ğ": "g",
				"ġ": "g",
				"ģ": "g",
				"Ĥ": "H",
				"Ħ": "H",
				"ĥ": "h",
				"ħ": "h",
				"Ĩ": "I",
				"Ī": "I",
				"Ĭ": "I",
				"Į": "I",
				"İ": "I",
				"ĩ": "i",
				"ī": "i",
				"ĭ": "i",
				"į": "i",
				"ı": "i",
				"Ĵ": "J",
				"ĵ": "j",
				"Ķ": "K",
				"ķ": "k",
				"ĸ": "k",
				"Ĺ": "L",
				"Ļ": "L",
				"Ľ": "L",
				"Ŀ": "L",
				"Ł": "L",
				"ĺ": "l",
				"ļ": "l",
				"ľ": "l",
				"ŀ": "l",
				"ł": "l",
				"Ń": "N",
				"Ņ": "N",
				"Ň": "N",
				"Ŋ": "N",
				"ń": "n",
				"ņ": "n",
				"ň": "n",
				"ŋ": "n",
				"Ō": "O",
				"Ŏ": "O",
				"Ő": "O",
				"ō": "o",
				"ŏ": "o",
				"ő": "o",
				"Ŕ": "R",
				"Ŗ": "R",
				"Ř": "R",
				"ŕ": "r",
				"ŗ": "r",
				"ř": "r",
				"Ś": "S",
				"Ŝ": "S",
				"Ş": "S",
				"Š": "S",
				"ś": "s",
				"ŝ": "s",
				"ş": "s",
				"š": "s",
				"Ţ": "T",
				"Ť": "T",
				"Ŧ": "T",
				"ţ": "t",
				"ť": "t",
				"ŧ": "t",
				"Ũ": "U",
				"Ū": "U",
				"Ŭ": "U",
				"Ů": "U",
				"Ű": "U",
				"Ų": "U",
				"ũ": "u",
				"ū": "u",
				"ŭ": "u",
				"ů": "u",
				"ű": "u",
				"ų": "u",
				"Ŵ": "W",
				"ŵ": "w",
				"Ŷ": "Y",
				"ŷ": "y",
				"Ÿ": "Y",
				"Ź": "Z",
				"Ż": "Z",
				"Ž": "Z",
				"ź": "z",
				"ż": "z",
				"ž": "z",
				"Ĳ": "IJ",
				"ĳ": "ij",
				"Œ": "Oe",
				"œ": "oe",
				"ŉ": "'n",
				"ſ": "s"
			},
			m = /[\xc0-\xd6\xd8-\xf6\xf8-\xff\u0100-\u017f]/g,
			v = RegExp("[\\u0300-\\u036f\\ufe20-\\ufe2f\\u20d0-\\u20ff\\u1ab0-\\u1aff\\u1dc0-\\u1dff]", "g");

		function b(t) {
			return f[t]
		}

		function y(t) {
			return (t = t.toString()) && t.replace(m, b).replace(v, "")
		}
		var w = {
			"&": "&amp;",
			"<": "&lt;",
			">": "&gt;",
			'"': "&quot;",
			"'": "&#x27;"
		};
		w[String.fromCharCode(96)] = "&#x60;";
		var x, E, D, C, T = (x = w, E = "(?:" + Object.keys(x).join("|") + ")", D = RegExp(E), C = RegExp(E, "g"), function(t) {
			return t = null == t ? "" : "" + t, D.test(t) ? t.replace(C, O) : t
		});

		function O(t) {
			return x[t]
		}
		var M = {
				32: " ",
				48: "0",
				49: "1",
				50: "2",
				51: "3",
				52: "4",
				53: "5",
				54: "6",
				55: "7",
				56: "8",
				57: "9",
				59: ";",
				65: "A",
				66: "B",
				67: "C",
				68: "D",
				69: "E",
				70: "F",
				71: "G",
				72: "H",
				73: "I",
				74: "J",
				75: "K",
				76: "L",
				77: "M",
				78: "N",
				79: "O",
				80: "P",
				81: "Q",
				82: "R",
				83: "S",
				84: "T",
				85: "U",
				86: "V",
				87: "W",
				88: "X",
				89: "Y",
				90: "Z",
				96: "0",
				97: "1",
				98: "2",
				99: "3",
				100: "4",
				101: "5",
				102: "6",
				103: "7",
				104: "8",
				105: "9"
			},
			N = 27,
			_ = 13,
			L = 32,
			P = 9,
			H = 38,
			F = 40,
			B = {
				success: !1,
				major: "3"
			};
		try {
			B.full = (k.fn.dropdown.Constructor.VERSION || "").split(" ")[0].split("."), B.major = B.full[0], B.success = !0
		} catch (t) {}
		var W = 0,
			j = ".bs.select",
			R = {
				DISABLED: "disabled",
				DIVIDER: "divider",
				SHOW: "open",
				DROPUP: "dropup",
				MENU: "dropdown-menu",
				MENURIGHT: "dropdown-menu-right",
				MENULEFT: "dropdown-menu-left",
				BUTTONCLASS: "btn-default",
				POPOVERHEADER: "popover-title",
				ICONBASE: "glyphicon",
				TICKICON: "glyphicon-ok"
			},
			U = {
				MENU: "." + R.MENU
			},
			V = {
				span: document.createElement("span"),
				i: document.createElement("i"),
				subtext: document.createElement("small"),
				a: document.createElement("a"),
				li: document.createElement("li"),
				whitespace: document.createTextNode(" "),
				fragment: document.createDocumentFragment()
			};
		V.a.setAttribute("role", "option"), V.subtext.className = "text-muted", V.text = V.span.cloneNode(!1), V.text.className = "text", V.checkMark = V.span.cloneNode(!1);
		var Y = new RegExp(H + "|" + F),
			X = new RegExp("^" + P + "$|" + N),
			K = function(t, e, n) {
				var i = V.li.cloneNode(!1);
				return t && (1 === t.nodeType || 11 === t.nodeType ? i.appendChild(t) : i.innerHTML = t), void 0 !== e && "" !== e && (i.className = e), null != n && i.classList.add("optgroup-" + n), i
			},
			G = function(t, e, n) {
				var i = V.a.cloneNode(!0);
				return t && (11 === t.nodeType ? i.appendChild(t) : i.insertAdjacentHTML("beforeend", t)), void 0 !== e && "" !== e && (i.className = e), "4" === B.major && i.classList.add("dropdown-item"), n && i.setAttribute("style", n), i
			},
			q = function(t, e) {
				var n, i, o = V.text.cloneNode(!1);
				if (t.content) o.innerHTML = t.content;
				else {
					if (o.textContent = t.text, t.icon) {
						var a = V.whitespace.cloneNode(!1);
						(i = (!0 === e ? V.i : V.span).cloneNode(!1)).className = t.iconBase + " " + t.icon, V.fragment.appendChild(i), V.fragment.appendChild(a)
					}
					t.subtext && ((n = V.subtext.cloneNode(!1)).textContent = t.subtext, o.appendChild(n))
				}
				if (!0 === e)
					for (; 0 < o.childNodes.length;) V.fragment.appendChild(o.childNodes[0]);
				else V.fragment.appendChild(o);
				return V.fragment
			},
			J = function(t) {
				var e, n, i = V.text.cloneNode(!1);
				if (i.innerHTML = t.label, t.icon) {
					var o = V.whitespace.cloneNode(!1);
					(n = V.span.cloneNode(!1)).className = t.iconBase + " " + t.icon, V.fragment.appendChild(n), V.fragment.appendChild(o)
				}
				return t.subtext && ((e = V.subtext.cloneNode(!1)).textContent = t.subtext, i.appendChild(e)), V.fragment.appendChild(i), V.fragment
			},
			Z = function(t, e) {
				var n = this;
				h.useDefault || (k.valHooks.select.set = h._set, h.useDefault = !0), this.$element = k(t), this.$newElement = null, this.$button = null, this.$menu = null, this.options = e, this.selectpicker = {
					main: {},
					current: {},
					search: {},
					view: {},
					keydown: {
						keyHistory: "",
						resetKeyHistory: {
							start: function() {
								return setTimeout(function() {
									n.selectpicker.keydown.keyHistory = ""
								}, 800)
							}
						}
					}
				}, null === this.options.title && (this.options.title = this.$element.attr("title"));
				var i = this.options.windowPadding;
				"number" == typeof i && (this.options.windowPadding = [i, i, i, i]), this.val = Z.prototype.val, this.render = Z.prototype.render, this.refresh = Z.prototype.refresh, this.setStyle = Z.prototype.setStyle, this.selectAll = Z.prototype.selectAll, this.deselectAll = Z.prototype.deselectAll, this.destroy = Z.prototype.destroy, this.remove = Z.prototype.remove, this.show = Z.prototype.show, this.hide = Z.prototype.hide, this.init()
			};

		function Q(t) {
			var r, l = arguments,
				c = t;
			if ([].shift.apply(l), !B.success) {
				try {
					B.full = (k.fn.dropdown.Constructor.VERSION || "").split(" ")[0].split(".")
				} catch (t) {
					Z.BootstrapVersion ? B.full = Z.BootstrapVersion.split(" ")[0].split(".") : (B.full = [B.major, "0", "0"], console.warn("There was an issue retrieving Bootstrap's version. Ensure Bootstrap is being loaded before bootstrap-select and there is no namespace collision. If loading Bootstrap asynchronously, the version may need to be manually specified via $.fn.selectpicker.Constructor.BootstrapVersion.", t))
				}
				B.major = B.full[0], B.success = !0
			}
			if ("4" === B.major) {
				var e = [];
				Z.DEFAULTS.style === R.BUTTONCLASS && e.push({
					name: "style",
					className: "BUTTONCLASS"
				}), Z.DEFAULTS.iconBase === R.ICONBASE && e.push({
					name: "iconBase",
					className: "ICONBASE"
				}), Z.DEFAULTS.tickIcon === R.TICKICON && e.push({
					name: "tickIcon",
					className: "TICKICON"
				}), R.DIVIDER = "dropdown-divider", R.SHOW = "show", R.BUTTONCLASS = "btn-light", R.POPOVERHEADER = "popover-header", R.ICONBASE = "", R.TICKICON = "bs-ok-default";
				for (var n = 0; n < e.length; n++) {
					t = e[n];
					Z.DEFAULTS[t.name] = R[t.className]
				}
			}
			var i = this.each(function() {
				var t = k(this);
				if (t.is("select")) {
					var e = t.data("selectpicker"),
						n = "object" == typeof c && c;
					if (e) {
						if (n)
							for (var i in n) n.hasOwnProperty(i) && (e.options[i] = n[i])
					} else {
						var o = t.data();
						for (var a in o) o.hasOwnProperty(a) && -1 !== k.inArray(a, d) && delete o[a];
						var s = k.extend({}, Z.DEFAULTS, k.fn.selectpicker.defaults || {}, o, n);
						s.template = k.extend({}, Z.DEFAULTS.template, k.fn.selectpicker.defaults ? k.fn.selectpicker.defaults.template : {}, o.template, n.template), t.data("selectpicker", e = new Z(this, s))
					}
					"string" == typeof c && (r = e[c] instanceof Function ? e[c].apply(e, l) : e.options[c])
				}
			});
			return void 0 !== r ? r : i
		}
		Z.VERSION = "1.13.9", Z.DEFAULTS = {
			noneSelectedText: "Nothing selected",
			noneResultsText: "No results matched {0}",
			countSelectedText: function(t, e) {
				return 1 == t ? "{0} item selected" : "{0} items selected"
			},
			maxOptionsText: function(t, e) {
				return [1 == t ? "Limit reached ({n} item max)" : "Limit reached ({n} items max)", 1 == e ? "Group limit reached ({n} item max)" : "Group limit reached ({n} items max)"]
			},
			selectAllText: "Select All",
			deselectAllText: "Deselect All",
			doneButton: !1,
			doneButtonText: "Close",
			multipleSeparator: ", ",
			styleBase: "btn",
			style: R.BUTTONCLASS,
			size: "auto",
			title: null,
			selectedTextFormat: "values",
			width: !1,
			container: !1,
			hideDisabled: !1,
			showSubtext: !1,
			showIcon: !0,
			showContent: !0,
			dropupAuto: !0,
			header: !1,
			liveSearch: !1,
			liveSearchPlaceholder: null,
			liveSearchNormalize: !1,
			liveSearchStyle: "contains",
			actionsBox: !1,
			iconBase: R.ICONBASE,
			tickIcon: R.TICKICON,
			showTick: !1,
			template: {
				caret: '<span class="caret"></span>'
			},
			maxOptions: !1,
			mobile: !1,
			selectOnTab: !1,
			dropdownAlignRight: !1,
			windowPadding: 0,
			virtualScroll: 600,
			display: !1,
			sanitize: !0,
			sanitizeFn: null,
			whiteList: t
		}, Z.prototype = {
			constructor: Z,
			init: function() {
				var n = this,
					t = this.$element.attr("id");
				this.selectId = W++, this.$element[0].classList.add("bs-select-hidden"), this.multiple = this.$element.prop("multiple"), this.autofocus = this.$element.prop("autofocus"), this.options.showTick = this.$element[0].classList.contains("show-tick"), this.$newElement = this.createDropdown(), this.$element.after(this.$newElement).prependTo(this.$newElement), this.$button = this.$newElement.children("button"), this.$menu = this.$newElement.children(U.MENU), this.$menuInner = this.$menu.children(".inner"), this.$searchbox = this.$menu.find("input"), this.$element[0].classList.remove("bs-select-hidden"), !0 === this.options.dropdownAlignRight && this.$menu[0].classList.add(R.MENURIGHT), void 0 !== t && this.$button.attr("data-id", t), this.checkDisabled(), this.clickListener(), this.options.liveSearch && this.liveSearchListener(), this.setStyle(), this.render(), this.setWidth(), this.options.container ? this.selectPosition() : this.$element.on("hide" + j, function() {
					if (n.isVirtual()) {
						var t = n.$menuInner[0],
							e = t.firstChild.cloneNode(!1);
						t.replaceChild(e, t.firstChild), t.scrollTop = 0
					}
				}), this.$menu.data("this", this), this.$newElement.data("this", this), this.options.mobile && this.mobile(), this.$newElement.on({
					"hide.bs.dropdown": function(t) {
						n.$menuInner.attr("aria-expanded", !1), n.$element.trigger("hide" + j, t)
					},
					"hidden.bs.dropdown": function(t) {
						n.$element.trigger("hidden" + j, t)
					},
					"show.bs.dropdown": function(t) {
						n.$menuInner.attr("aria-expanded", !0), n.$element.trigger("show" + j, t)
					},
					"shown.bs.dropdown": function(t) {
						n.$element.trigger("shown" + j, t)
					}
				}), n.$element[0].hasAttribute("required") && this.$element.on("invalid" + j, function() {
					n.$button[0].classList.add("bs-invalid"), n.$element.on("shown" + j + ".invalid", function() {
						n.$element.val(n.$element.val()).off("shown" + j + ".invalid")
					}).on("rendered" + j, function() {
						this.validity.valid && n.$button[0].classList.remove("bs-invalid"), n.$element.off("rendered" + j)
					}), n.$button.on("blur" + j, function() {
						n.$element.trigger("focus").trigger("blur"), n.$button.off("blur" + j)
					})
				}), setTimeout(function() {
					n.createLi(), n.$element.trigger("loaded" + j)
				})
			},
			createDropdown: function() {
				var t = this.multiple || this.options.showTick ? " show-tick" : "",
					e = "",
					n = this.autofocus ? " autofocus" : "";
				B.major < 4 && this.$element.parent().hasClass("input-group") && (e = " input-group-btn");
				var i, o = "",
					a = "",
					s = "",
					r = "";
				return this.options.header && (o = '<div class="' + R.POPOVERHEADER + '"><button type="button" class="close" aria-hidden="true">&times;</button>' + this.options.header + "</div>"), this.options.liveSearch && (a = '<div class="bs-searchbox"><input type="text" class="form-control" autocomplete="off"' + (null === this.options.liveSearchPlaceholder ? "" : ' placeholder="' + T(this.options.liveSearchPlaceholder) + '"') + ' role="textbox" aria-label="Search"></div>'), this.multiple && this.options.actionsBox && (s = '<div class="bs-actionsbox"><div class="btn-group btn-group-sm btn-block"><button type="button" class="actions-btn bs-select-all btn ' + R.BUTTONCLASS + '">' + this.options.selectAllText + '</button><button type="button" class="actions-btn bs-deselect-all btn ' + R.BUTTONCLASS + '">' + this.options.deselectAllText + "</button></div></div>"), this.multiple && this.options.doneButton && (r = '<div class="bs-donebutton"><div class="btn-group btn-block"><button type="button" class="btn btn-sm ' + R.BUTTONCLASS + '">' + this.options.doneButtonText + "</button></div></div>"), i = '<div class="dropdown bootstrap-select' + t + e + '"><button type="button" class="' + this.options.styleBase + ' dropdown-toggle" ' + ("static" === this.options.display ? 'data-display="static"' : "") + 'data-toggle="dropdown"' + n + ' role="button"><div class="filter-option"><div class="filter-option-inner"><div class="filter-option-inner-inner"></div></div> </div>' + ("4" === B.major ? "" : '<span class="bs-caret">' + this.options.template.caret + "</span>") + '</button><div class="' + R.MENU + " " + ("4" === B.major ? "" : R.SHOW) + '" role="combobox">' + o + a + s + '<div class="inner ' + R.SHOW + '" role="listbox" aria-expanded="false" tabindex="-1"><ul class="' + R.MENU + " inner " + ("4" === B.major ? R.SHOW : "") + '"></ul></div>' + r + "</div></div>", k(i)
			},
			setPositionData: function() {
				this.selectpicker.view.canHighlight = [];
				for (var t = 0; t < this.selectpicker.current.data.length; t++) {
					var e = this.selectpicker.current.data[t],
						n = !0;
					"divider" === e.type ? (n = !1, e.height = this.sizeInfo.dividerHeight) : "optgroup-label" === e.type ? (n = !1, e.height = this.sizeInfo.dropdownHeaderHeight) : e.height = this.sizeInfo.liHeight, e.disabled && (n = !1), this.selectpicker.view.canHighlight.push(n), e.position = (0 === t ? 0 : this.selectpicker.current.data[t - 1].position) + e.height
				}
			},
			isVirtual: function() {
				return !1 !== this.options.virtualScroll && this.selectpicker.main.elements.length >= this.options.virtualScroll || !0 === this.options.virtualScroll
			},
			createView: function(O, t) {
				t = t || 0;
				var M = this;
				this.selectpicker.current = O ? this.selectpicker.search : this.selectpicker.main;
				var N, _, L = [];

				function n(t, e) {
					var n, i, o, a, s, r, l, c, d, u, h = M.selectpicker.current.elements.length,
						p = [],
						f = !0,
						m = M.isVirtual();
					M.selectpicker.view.scrollTop = t, !0 === m && M.sizeInfo.hasScrollBar && M.$menu[0].offsetWidth > M.sizeInfo.totalMenuWidth && (M.sizeInfo.menuWidth = M.$menu[0].offsetWidth, M.sizeInfo.totalMenuWidth = M.sizeInfo.menuWidth + M.sizeInfo.scrollBarWidth, M.$menu.css("min-width", M.sizeInfo.menuWidth)), n = Math.ceil(M.sizeInfo.menuInnerHeight / M.sizeInfo.liHeight * 1.5), i = Math.round(h / n) || 1;
					for (var g = 0; g < i; g++) {
						var v = (g + 1) * n;
						if (g === i - 1 && (v = h), p[g] = [g * n + (g ? 1 : 0), v], !h) break;
						void 0 === s && t <= M.selectpicker.current.data[v - 1].position - M.sizeInfo.menuInnerHeight && (s = g)
					}
					if (void 0 === s && (s = 0), r = [M.selectpicker.view.position0, M.selectpicker.view.position1], o = Math.max(0, s - 1), a = Math.min(i - 1, s + 1), M.selectpicker.view.position0 = !1 === m ? 0 : Math.max(0, p[o][0]) || 0, M.selectpicker.view.position1 = !1 === m ? h : Math.min(h, p[a][1]) || 0, l = r[0] !== M.selectpicker.view.position0 || r[1] !== M.selectpicker.view.position1, void 0 !== M.activeIndex && (_ = M.selectpicker.main.elements[M.prevActiveIndex], L = M.selectpicker.main.elements[M.activeIndex], N = M.selectpicker.main.elements[M.selectedIndex], e && (M.activeIndex !== M.selectedIndex && L && L.length && (L.classList.remove("active"), L.firstChild && L.firstChild.classList.remove("active")), M.activeIndex = void 0), M.activeIndex && M.activeIndex !== M.selectedIndex && N && N.length && (N.classList.remove("active"), N.firstChild && N.firstChild.classList.remove("active"))), void 0 !== M.prevActiveIndex && M.prevActiveIndex !== M.activeIndex && M.prevActiveIndex !== M.selectedIndex && _ && _.length && (_.classList.remove("active"), _.firstChild && _.firstChild.classList.remove("active")), (e || l) && (c = M.selectpicker.view.visibleElements ? M.selectpicker.view.visibleElements.slice() : [], M.selectpicker.view.visibleElements = !1 === m ? M.selectpicker.current.elements : M.selectpicker.current.elements.slice(M.selectpicker.view.position0, M.selectpicker.view.position1), M.setOptionStatus(), (O || !1 === m && e) && (d = c, u = M.selectpicker.view.visibleElements, f = !(d.length === u.length && d.every(function(t, e) {
							return t === u[e]
						}))), (e || !0 === m) && f)) {
						var b, y, w = M.$menuInner[0],
							x = document.createDocumentFragment(),
							E = w.firstChild.cloneNode(!1),
							S = M.selectpicker.view.visibleElements,
							D = [];
						w.replaceChild(E, w.firstChild);
						g = 0;
						for (var C = S.length; g < C; g++) {
							var I, $, T = S[g];
							M.options.sanitize && (I = T.lastChild) && ($ = M.selectpicker.current.data[g + M.selectpicker.view.position0]) && $.content && !$.sanitized && (D.push(I), $.sanitized = !0), x.appendChild(T)
						}
						M.options.sanitize && D.length && z(D, M.options.whiteList, M.options.sanitizeFn), !0 === m && (b = 0 === M.selectpicker.view.position0 ? 0 : M.selectpicker.current.data[M.selectpicker.view.position0 - 1].position, y = M.selectpicker.view.position1 > h - 1 ? 0 : M.selectpicker.current.data[h - 1].position - M.selectpicker.current.data[M.selectpicker.view.position1 - 1].position, w.firstChild.style.marginTop = b + "px", w.firstChild.style.marginBottom = y + "px"), w.firstChild.appendChild(x)
					}
					if (M.prevActiveIndex = M.activeIndex, M.options.liveSearch) {
						if (O && e) {
							var k, A = 0;
							M.selectpicker.view.canHighlight[A] || (A = 1 + M.selectpicker.view.canHighlight.slice(1).indexOf(!0)), k = M.selectpicker.view.visibleElements[A], M.selectpicker.view.currentActive && (M.selectpicker.view.currentActive.classList.remove("active"), M.selectpicker.view.currentActive.firstChild && M.selectpicker.view.currentActive.firstChild.classList.remove("active")), k && (k.classList.add("active"), k.firstChild && k.firstChild.classList.add("active")), M.activeIndex = (M.selectpicker.current.data[A] || {}).index
						}
					} else M.$menuInner.trigger("focus")
				}
				this.setPositionData(), n(t, !0), this.$menuInner.off("scroll.createView").on("scroll.createView", function(t, e) {
					M.noScroll || n(this.scrollTop, e), M.noScroll = !1
				}), k(window).off("resize" + j + "." + this.selectId + ".createView").on("resize" + j + "." + this.selectId + ".createView", function() {
					M.$newElement.hasClass(R.SHOW) && n(M.$menuInner[0].scrollTop)
				})
			},
			setPlaceholder: function() {
				var t = !1;
				if (this.options.title && !this.multiple) {
					this.selectpicker.view.titleOption || (this.selectpicker.view.titleOption = document.createElement("option")), t = !0;
					var e = this.$element[0],
						n = !1,
						i = !this.selectpicker.view.titleOption.parentNode;
					if (i) this.selectpicker.view.titleOption.className = "bs-title-option", this.selectpicker.view.titleOption.value = "", n = void 0 === k(e.options[e.selectedIndex]).attr("selected") && void 0 === this.$element.data("selected");
					!i && 0 === this.selectpicker.view.titleOption.index || e.insertBefore(this.selectpicker.view.titleOption, e.firstChild), n && (e.selectedIndex = 0)
				}
				return t
			},
			createLi: function() {
				var l = this,
					f = this.options.iconBase,
					m = ':not([hidden]):not([data-hidden="true"])',
					g = [],
					v = [],
					c = 0,
					b = 0,
					t = this.setPlaceholder() ? 1 : 0;
				this.options.hideDisabled && (m += ":not(:disabled)"), !l.options.showTick && !l.multiple || V.checkMark.parentNode || (V.checkMark.className = f + " " + l.options.tickIcon + " check-mark", V.a.appendChild(V.checkMark));
				var e = this.$element[0].querySelectorAll("select > *" + m);

				function y(t) {
					var e = v[v.length - 1];
					e && "divider" === e.type && (e.optID || t.optID) || ((t = t || {}).type = "divider", g.push(K(!1, R.DIVIDER, t.optID ? t.optID + "div" : void 0)), v.push(t))
				}

				function w(t, e) {
					if ((e = e || {}).divider = "true" === t.getAttribute("data-divider"), e.divider) y({
						optID: e.optID
					});
					else {
						var n = v.length,
							i = t.style.cssText,
							o = i ? T(i) : "",
							a = (t.className || "") + (e.optgroupClass || "");
						e.optID && (a = "opt " + a), e.text = t.textContent, e.content = t.getAttribute("data-content"), e.tokens = t.getAttribute("data-tokens"), e.subtext = t.getAttribute("data-subtext"), e.icon = t.getAttribute("data-icon"), e.iconBase = f;
						var s = q(e);
						g.push(K(G(s, a, o), "", e.optID)), t.liIndex = n, e.display = e.content || e.text, e.type = "option", e.index = n, e.option = t, e.disabled = e.disabled || t.disabled, v.push(e);
						var r = 0;
						e.display && (r += e.display.length), e.subtext && (r += e.subtext.length), e.icon && (r += 1), c < r && (c = r, l.selectpicker.view.widestOption = g[g.length - 1])
					}
				}

				function n(t, e) {
					var n = e[t],
						i = e[t - 1],
						o = e[t + 1],
						a = n.querySelectorAll("option" + m);
					if (a.length) {
						var s, r, l = {
								label: T(n.label),
								subtext: n.getAttribute("data-subtext"),
								icon: n.getAttribute("data-icon"),
								iconBase: f
							},
							c = " " + (n.className || "");
						b++, i && y({
							optID: b
						});
						var d = J(l);
						g.push(K(d, "dropdown-header" + c, b)), v.push({
							display: l.label,
							subtext: l.subtext,
							type: "optgroup-label",
							optID: b
						});
						for (var u = 0, h = a.length; u < h; u++) {
							var p = a[u];
							0 === u && (r = (s = v.length - 1) + h), w(p, {
								headerIndex: s,
								lastIndex: r,
								optID: b,
								optgroupClass: c,
								disabled: n.disabled
							})
						}
						o && y({
							optID: b
						})
					}
				}
				for (var i = e.length; t < i; t++) {
					var o = e[t];
					"OPTGROUP" !== o.tagName ? w(o, {}) : n(t, e)
				}
				this.selectpicker.main.elements = g, this.selectpicker.main.data = v, this.selectpicker.current = this.selectpicker.main
			},
			findLis: function() {
				return this.$menuInner.find(".inner > li")
			},
			render: function() {
				this.setPlaceholder();
				var t, e = this,
					n = this.$element[0].selectedOptions,
					i = n.length,
					o = this.$button[0],
					a = o.querySelector(".filter-option-inner-inner"),
					s = document.createTextNode(this.options.multipleSeparator),
					r = V.fragment.cloneNode(!1),
					l = !1;
				if (this.togglePlaceholder(), this.tabIndex(), "static" === this.options.selectedTextFormat) r = q({
					text: this.options.title
				}, !0);
				else if (!1 === (this.multiple && -1 !== this.options.selectedTextFormat.indexOf("count") && 1 < i && (1 < (t = this.options.selectedTextFormat.split(">")).length && i > t[1] || 1 === t.length && 2 <= i))) {
					for (var c = 0; c < i && c < 50; c++) {
						var d = n[c],
							u = {},
							h = {
								content: d.getAttribute("data-content"),
								subtext: d.getAttribute("data-subtext"),
								icon: d.getAttribute("data-icon")
							};
						this.multiple && 0 < c && r.appendChild(s.cloneNode(!1)), d.title ? u.text = d.title : h.content && e.options.showContent ? (u.content = h.content.toString(), l = !0) : (e.options.showIcon && (u.icon = h.icon, u.iconBase = this.options.iconBase), e.options.showSubtext && !e.multiple && h.subtext && (u.subtext = " " + h.subtext), u.text = d.textContent.trim()), r.appendChild(q(u, !0))
					}
					49 < i && r.appendChild(document.createTextNode("..."))
				} else {
					var p = ':not([hidden]):not([data-hidden="true"]):not([data-divider="true"])';
					this.options.hideDisabled && (p += ":not(:disabled)");
					var f = this.$element[0].querySelectorAll("select > option" + p + ", optgroup" + p + " option" + p).length,
						m = "function" == typeof this.options.countSelectedText ? this.options.countSelectedText(i, f) : this.options.countSelectedText;
					r = q({
						text: m.replace("{0}", i.toString()).replace("{1}", f.toString())
					}, !0)
				}
				if (null == this.options.title && (this.options.title = this.$element.attr("title")), r.childNodes.length || (r = q({
						text: void 0 !== this.options.title ? this.options.title : this.options.noneSelectedText
					}, !0)), o.title = r.textContent.replace(/<[^>]*>?/g, "").trim(), this.options.sanitize && l && z([r], e.options.whiteList, e.options.sanitizeFn), a.innerHTML = "", a.appendChild(r), B.major < 4 && this.$newElement[0].classList.contains("bs3-has-addon")) {
					var g = o.querySelector(".filter-expand"),
						v = a.cloneNode(!0);
					v.className = "filter-expand", g ? o.replaceChild(v, g) : o.appendChild(v)
				}
				this.$element.trigger("rendered" + j)
			},
			setStyle: function(t, e) {
				var n, i = this.$button[0],
					o = this.$newElement[0],
					a = this.options.style.trim();
				this.$element.attr("class") && this.$newElement.addClass(this.$element.attr("class").replace(/selectpicker|mobile-device|bs-select-hidden|validate\[.*\]/gi, "")), B.major < 4 && (o.classList.add("bs3"), o.parentNode.classList.contains("input-group") && (o.previousElementSibling || o.nextElementSibling) && (o.previousElementSibling || o.nextElementSibling).classList.contains("input-group-addon") && o.classList.add("bs3-has-addon")), n = t ? t.trim() : a, "add" == e ? n && i.classList.add.apply(i.classList, n.split(" ")) : "remove" == e ? n && i.classList.remove.apply(i.classList, n.split(" ")) : (a && i.classList.remove.apply(i.classList, a.split(" ")), n && i.classList.add.apply(i.classList, n.split(" ")))
			},
			liHeight: function(t) {
				if (t || !1 !== this.options.size && !this.sizeInfo) {
					this.sizeInfo || (this.sizeInfo = {});
					var e = document.createElement("div"),
						n = document.createElement("div"),
						i = document.createElement("div"),
						o = document.createElement("ul"),
						a = document.createElement("li"),
						s = document.createElement("li"),
						r = document.createElement("li"),
						l = document.createElement("a"),
						c = document.createElement("span"),
						d = this.options.header && 0 < this.$menu.find("." + R.POPOVERHEADER).length ? this.$menu.find("." + R.POPOVERHEADER)[0].cloneNode(!0) : null,
						u = this.options.liveSearch ? document.createElement("div") : null,
						h = this.options.actionsBox && this.multiple && 0 < this.$menu.find(".bs-actionsbox").length ? this.$menu.find(".bs-actionsbox")[0].cloneNode(!0) : null,
						p = this.options.doneButton && this.multiple && 0 < this.$menu.find(".bs-donebutton").length ? this.$menu.find(".bs-donebutton")[0].cloneNode(!0) : null,
						f = this.$element.find("option")[0];
					if (this.sizeInfo.selectWidth = this.$newElement[0].offsetWidth, c.className = "text", l.className = "dropdown-item " + (f ? f.className : ""), e.className = this.$menu[0].parentNode.className + " " + R.SHOW, e.style.width = this.sizeInfo.selectWidth + "px", "auto" === this.options.width && (n.style.minWidth = 0), n.className = R.MENU + " " + R.SHOW, i.className = "inner " + R.SHOW, o.className = R.MENU + " inner " + ("4" === B.major ? R.SHOW : ""), a.className = R.DIVIDER, s.className = "dropdown-header", c.appendChild(document.createTextNode("​")), l.appendChild(c), r.appendChild(l), s.appendChild(c.cloneNode(!0)), this.selectpicker.view.widestOption && o.appendChild(this.selectpicker.view.widestOption.cloneNode(!0)), o.appendChild(r), o.appendChild(a), o.appendChild(s), d && n.appendChild(d), u) {
						var m = document.createElement("input");
						u.className = "bs-searchbox", m.className = "form-control", u.appendChild(m), n.appendChild(u)
					}
					h && n.appendChild(h), i.appendChild(o), n.appendChild(i), p && n.appendChild(p), e.appendChild(n), document.body.appendChild(e);
					var g, v = r.offsetHeight,
						b = s ? s.offsetHeight : 0,
						y = d ? d.offsetHeight : 0,
						w = u ? u.offsetHeight : 0,
						x = h ? h.offsetHeight : 0,
						E = p ? p.offsetHeight : 0,
						S = k(a).outerHeight(!0),
						D = !!window.getComputedStyle && window.getComputedStyle(n),
						C = n.offsetWidth,
						I = D ? null : k(n),
						$ = {
							vert: A(D ? D.paddingTop : I.css("paddingTop")) + A(D ? D.paddingBottom : I.css("paddingBottom")) + A(D ? D.borderTopWidth : I.css("borderTopWidth")) + A(D ? D.borderBottomWidth : I.css("borderBottomWidth")),
							horiz: A(D ? D.paddingLeft : I.css("paddingLeft")) + A(D ? D.paddingRight : I.css("paddingRight")) + A(D ? D.borderLeftWidth : I.css("borderLeftWidth")) + A(D ? D.borderRightWidth : I.css("borderRightWidth"))
						},
						T = {
							vert: $.vert + A(D ? D.marginTop : I.css("marginTop")) + A(D ? D.marginBottom : I.css("marginBottom")) + 2,
							horiz: $.horiz + A(D ? D.marginLeft : I.css("marginLeft")) + A(D ? D.marginRight : I.css("marginRight")) + 2
						};
					i.style.overflowY = "scroll", g = n.offsetWidth - C, document.body.removeChild(e), this.sizeInfo.liHeight = v, this.sizeInfo.dropdownHeaderHeight = b, this.sizeInfo.headerHeight = y, this.sizeInfo.searchHeight = w, this.sizeInfo.actionsHeight = x, this.sizeInfo.doneButtonHeight = E, this.sizeInfo.dividerHeight = S, this.sizeInfo.menuPadding = $, this.sizeInfo.menuExtras = T, this.sizeInfo.menuWidth = C, this.sizeInfo.totalMenuWidth = this.sizeInfo.menuWidth, this.sizeInfo.scrollBarWidth = g, this.sizeInfo.selectHeight = this.$newElement[0].offsetHeight, this.setPositionData()
				}
			},
			getSelectPosition: function() {
				var t, e = k(window),
					n = this.$newElement.offset(),
					i = k(this.options.container);
				this.options.container && i.length && !i.is("body") ? ((t = i.offset()).top += parseInt(i.css("borderTopWidth")), t.left += parseInt(i.css("borderLeftWidth"))) : t = {
					top: 0,
					left: 0
				};
				var o = this.options.windowPadding;
				this.sizeInfo.selectOffsetTop = n.top - t.top - e.scrollTop(), this.sizeInfo.selectOffsetBot = e.height() - this.sizeInfo.selectOffsetTop - this.sizeInfo.selectHeight - t.top - o[2], this.sizeInfo.selectOffsetLeft = n.left - t.left - e.scrollLeft(), this.sizeInfo.selectOffsetRight = e.width() - this.sizeInfo.selectOffsetLeft - this.sizeInfo.selectWidth - t.left - o[1], this.sizeInfo.selectOffsetTop -= o[0], this.sizeInfo.selectOffsetLeft -= o[3]
			},
			setMenuSize: function(t) {
				this.getSelectPosition();
				var e, n, i, o, a, s, r, l = this.sizeInfo.selectWidth,
					c = this.sizeInfo.liHeight,
					d = this.sizeInfo.headerHeight,
					u = this.sizeInfo.searchHeight,
					h = this.sizeInfo.actionsHeight,
					p = this.sizeInfo.doneButtonHeight,
					f = this.sizeInfo.dividerHeight,
					m = this.sizeInfo.menuPadding,
					g = 0;
				if (this.options.dropupAuto && (r = c * this.selectpicker.current.elements.length + m.vert, this.$newElement.toggleClass(R.DROPUP, this.sizeInfo.selectOffsetTop - this.sizeInfo.selectOffsetBot > this.sizeInfo.menuExtras.vert && r + this.sizeInfo.menuExtras.vert + 50 > this.sizeInfo.selectOffsetBot)), "auto" === this.options.size) o = 3 < this.selectpicker.current.elements.length ? 3 * this.sizeInfo.liHeight + this.sizeInfo.menuExtras.vert - 2 : 0, n = this.sizeInfo.selectOffsetBot - this.sizeInfo.menuExtras.vert, i = o + d + u + h + p, s = Math.max(o - m.vert, 0), this.$newElement.hasClass(R.DROPUP) && (n = this.sizeInfo.selectOffsetTop - this.sizeInfo.menuExtras.vert), e = (a = n) - d - u - h - p - m.vert;
				else if (this.options.size && "auto" != this.options.size && this.selectpicker.current.elements.length > this.options.size) {
					for (var v = 0; v < this.options.size; v++) "divider" === this.selectpicker.current.data[v].type && g++;
					e = (n = c * this.options.size + g * f + m.vert) - m.vert, a = n + d + u + h + p, i = s = ""
				}
				"auto" === this.options.dropdownAlignRight && this.$menu.toggleClass(R.MENURIGHT, this.sizeInfo.selectOffsetLeft > this.sizeInfo.selectOffsetRight && this.sizeInfo.selectOffsetRight < this.sizeInfo.totalMenuWidth - l), this.$menu.css({
					"max-height": a + "px",
					overflow: "hidden",
					"min-height": i + "px"
				}), this.$menuInner.css({
					"max-height": e + "px",
					"overflow-y": "auto",
					"min-height": s + "px"
				}), this.sizeInfo.menuInnerHeight = Math.max(e, 1), this.selectpicker.current.data.length && this.selectpicker.current.data[this.selectpicker.current.data.length - 1].position > this.sizeInfo.menuInnerHeight && (this.sizeInfo.hasScrollBar = !0, this.sizeInfo.totalMenuWidth = this.sizeInfo.menuWidth + this.sizeInfo.scrollBarWidth, this.$menu.css("min-width", this.sizeInfo.totalMenuWidth)), this.dropdown && this.dropdown._popper && this.dropdown._popper.update()
			},
			setSize: function(t) {
				if (this.liHeight(t), this.options.header && this.$menu.css("padding-top", 0), !1 !== this.options.size) {
					var e, n = this,
						i = k(window),
						o = 0;
					if (this.setMenuSize(), this.options.liveSearch && this.$searchbox.off("input.setMenuSize propertychange.setMenuSize").on("input.setMenuSize propertychange.setMenuSize", function() {
							return n.setMenuSize()
						}), "auto" === this.options.size ? i.off("resize" + j + "." + this.selectId + ".setMenuSize scroll" + j + "." + this.selectId + ".setMenuSize").on("resize" + j + "." + this.selectId + ".setMenuSize scroll" + j + "." + this.selectId + ".setMenuSize", function() {
							return n.setMenuSize()
						}) : this.options.size && "auto" != this.options.size && this.selectpicker.current.elements.length > this.options.size && i.off("resize" + j + "." + this.selectId + ".setMenuSize scroll" + j + "." + this.selectId + ".setMenuSize"), t) o = this.$menuInner[0].scrollTop;
					else if (!n.multiple) {
						var a = n.$element[0];
						"number" == typeof(e = (a.options[a.selectedIndex] || {}).liIndex) && !1 !== n.options.size && (o = (o = n.sizeInfo.liHeight * e) - n.sizeInfo.menuInnerHeight / 2 + n.sizeInfo.liHeight / 2)
					}
					n.createView(!1, o)
				}
			},
			setWidth: function() {
				var n = this;
				"auto" === this.options.width ? requestAnimationFrame(function() {
					n.$menu.css("min-width", "0"), n.$element.on("loaded" + j, function() {
						n.liHeight(), n.setMenuSize();
						var t = n.$newElement.clone().appendTo("body"),
							e = t.css("width", "auto").children("button").outerWidth();
						t.remove(), n.sizeInfo.selectWidth = Math.max(n.sizeInfo.totalMenuWidth, e), n.$newElement.css("width", n.sizeInfo.selectWidth + "px")
					})
				}) : "fit" === this.options.width ? (this.$menu.css("min-width", ""), this.$newElement.css("width", "").addClass("fit-width")) : this.options.width ? (this.$menu.css("min-width", ""), this.$newElement.css("width", this.options.width)) : (this.$menu.css("min-width", ""), this.$newElement.css("width", "")), this.$newElement.hasClass("fit-width") && "fit" !== this.options.width && this.$newElement[0].classList.remove("fit-width")
			},
			selectPosition: function() {
				this.$bsContainer = k('<div class="bs-container" />');

				function t(t) {
					var e = {},
						n = s.options.display || !!k.fn.dropdown.Constructor.Default && k.fn.dropdown.Constructor.Default.display;
					s.$bsContainer.addClass(t.attr("class").replace(/form-control|fit-width/gi, "")).toggleClass(R.DROPUP, t.hasClass(R.DROPUP)), i = t.offset(), r.is("body") ? o = {
						top: 0,
						left: 0
					} : ((o = r.offset()).top += parseInt(r.css("borderTopWidth")) - r.scrollTop(), o.left += parseInt(r.css("borderLeftWidth")) - r.scrollLeft()), a = t.hasClass(R.DROPUP) ? 0 : t[0].offsetHeight, (B.major < 4 || "static" === n) && (e.top = i.top - o.top + a, e.left = i.left - o.left), e.width = t[0].offsetWidth, s.$bsContainer.css(e)
				}
				var i, o, a, s = this,
					r = k(this.options.container);
				this.$button.on("click.bs.dropdown.data-api", function() {
					s.isDisabled() || (t(s.$newElement), s.$bsContainer.appendTo(s.options.container).toggleClass(R.SHOW, !s.$button.hasClass(R.SHOW)).append(s.$menu))
				}), k(window).off("resize" + j + "." + this.selectId + " scroll" + j + "." + this.selectId).on("resize" + j + "." + this.selectId + " scroll" + j + "." + this.selectId, function() {
					s.$newElement.hasClass(R.SHOW) && t(s.$newElement)
				}), this.$element.on("hide" + j, function() {
					s.$menu.data("height", s.$menu.height()), s.$bsContainer.detach()
				})
			},
			setOptionStatus: function() {
				var t = this;
				if (t.noScroll = !1, t.selectpicker.view.visibleElements && t.selectpicker.view.visibleElements.length)
					for (var e = 0; e < t.selectpicker.view.visibleElements.length; e++) {
						var n = t.selectpicker.current.data[e + t.selectpicker.view.position0],
							i = n.option;
						i && (t.setDisabled(n.index, n.disabled), t.setSelected(n.index, i.selected))
					}
			},
			setSelected: function(t, e) {
				var n, i, o = this.selectpicker.main.elements[t],
					a = this.selectpicker.main.data[t],
					s = void 0 !== this.activeIndex,
					r = this.activeIndex === t || e && !this.multiple && !s;
				a.selected = e, i = o.firstChild, e && (this.selectedIndex = t), o.classList.toggle("selected", e), o.classList.toggle("active", r), r && (this.selectpicker.view.currentActive = o, this.activeIndex = t), i && (i.classList.toggle("selected", e), i.classList.toggle("active", r), i.setAttribute("aria-selected", e)), r || !s && e && void 0 !== this.prevActiveIndex && ((n = this.selectpicker.main.elements[this.prevActiveIndex]).classList.remove("active"), n.firstChild && n.firstChild.classList.remove("active"))
			},
			setDisabled: function(t, e) {
				var n, i = this.selectpicker.main.elements[t];
				this.selectpicker.main.data[t].disabled = e, n = i.firstChild, i.classList.toggle(R.DISABLED, e), n && ("4" === B.major && n.classList.toggle(R.DISABLED, e), n.setAttribute("aria-disabled", e), e ? n.setAttribute("tabindex", -1) : n.setAttribute("tabindex", 0))
			},
			isDisabled: function() {
				return this.$element[0].disabled
			},
			checkDisabled: function() {
				var t = this;
				this.isDisabled() ? (this.$newElement[0].classList.add(R.DISABLED), this.$button.addClass(R.DISABLED).attr("tabindex", -1).attr("aria-disabled", !0)) : (this.$button[0].classList.contains(R.DISABLED) && (this.$newElement[0].classList.remove(R.DISABLED), this.$button.removeClass(R.DISABLED).attr("aria-disabled", !1)), -1 != this.$button.attr("tabindex") || this.$element.data("tabindex") || this.$button.removeAttr("tabindex")), this.$button.on("click", function() {
					return !t.isDisabled()
				})
			},
			togglePlaceholder: function() {
				var t = this.$element[0],
					e = t.selectedIndex,
					n = -1 === e;
				n || t.options[e].value || (n = !0), this.$button.toggleClass("bs-placeholder", n)
			},
			tabIndex: function() {
				this.$element.data("tabindex") !== this.$element.attr("tabindex") && -98 !== this.$element.attr("tabindex") && "-98" !== this.$element.attr("tabindex") && (this.$element.data("tabindex", this.$element.attr("tabindex")), this.$button.attr("tabindex", this.$element.data("tabindex"))), this.$element.attr("tabindex", -98)
			},
			clickListener: function() {
				var C = this,
					e = k(document);

				function t() {
					C.options.liveSearch ? C.$searchbox.trigger("focus") : C.$menuInner.trigger("focus")
				}

				function n() {
					C.dropdown && C.dropdown._popper && C.dropdown._popper.state.isCreated ? t() : requestAnimationFrame(n)
				}
				e.data("spaceSelect", !1), this.$button.on("keyup", function(t) {
					/(32)/.test(t.keyCode.toString(10)) && e.data("spaceSelect") && (t.preventDefault(), e.data("spaceSelect", !1))
				}), this.$newElement.on("show.bs.dropdown", function() {
					3 < B.major && !C.dropdown && (C.dropdown = C.$button.data("bs.dropdown"), C.dropdown._menu = C.$menu[0])
				}), this.$button.on("click.bs.dropdown.data-api", function() {
					C.$newElement.hasClass(R.SHOW) || C.setSize()
				}), this.$element.on("shown" + j, function() {
					C.$menuInner[0].scrollTop !== C.selectpicker.view.scrollTop && (C.$menuInner[0].scrollTop = C.selectpicker.view.scrollTop), 3 < B.major ? requestAnimationFrame(n) : t()
				}), this.$menuInner.on("click", "li a", function(t, e) {
					var n = k(this),
						i = C.isVirtual() ? C.selectpicker.view.position0 : 0,
						o = C.selectpicker.current.data[n.parent().index() + i],
						a = o.index,
						s = I(C.$element[0]),
						r = C.$element.prop("selectedIndex"),
						l = !0;
					if (C.multiple && 1 !== C.options.maxOptions && t.stopPropagation(), t.preventDefault(), !C.isDisabled() && !n.parent().hasClass(R.DISABLED)) {
						var c = C.$element.find("option"),
							d = o.option,
							u = k(d),
							h = d.selected,
							p = u.parent("optgroup"),
							f = p.find("option"),
							m = C.options.maxOptions,
							g = p.data("maxOptions") || !1;
						if (a === C.activeIndex && (e = !0), e || (C.prevActiveIndex = C.activeIndex, C.activeIndex = void 0), C.multiple) {
							if (d.selected = !h, C.setSelected(a, !h), n.trigger("blur"), !1 !== m || !1 !== g) {
								var v = m < c.filter(":selected").length,
									b = g < p.find("option:selected").length;
								if (m && v || g && b)
									if (m && 1 == m) {
										c.prop("selected", !1), u.prop("selected", !0);
										for (var y = 0; y < c.length; y++) C.setSelected(y, !1);
										C.setSelected(a, !0)
									} else if (g && 1 == g) {
									p.find("option:selected").prop("selected", !1), u.prop("selected", !0);
									for (y = 0; y < f.length; y++) {
										d = f[y];
										C.setSelected(c.index(d), !1)
									}
									C.setSelected(a, !0)
								} else {
									var w = "string" == typeof C.options.maxOptionsText ? [C.options.maxOptionsText, C.options.maxOptionsText] : C.options.maxOptionsText,
										x = "function" == typeof w ? w(m, g) : w,
										E = x[0].replace("{n}", m),
										S = x[1].replace("{n}", g),
										D = k('<div class="notify"></div>');
									x[2] && (E = E.replace("{var}", x[2][1 < m ? 0 : 1]), S = S.replace("{var}", x[2][1 < g ? 0 : 1])), u.prop("selected", !1), C.$menu.append(D), m && v && (D.append(k("<div>" + E + "</div>")), l = !1, C.$element.trigger("maxReached" + j)), g && b && (D.append(k("<div>" + S + "</div>")), l = !1, C.$element.trigger("maxReachedGrp" + j)), setTimeout(function() {
										C.setSelected(a, !1)
									}, 10), D.delay(750).fadeOut(300, function() {
										k(this).remove()
									})
								}
							}
						} else c.prop("selected", !1), d.selected = !0, C.setSelected(a, !0);
						!C.multiple || C.multiple && 1 === C.options.maxOptions ? C.$button.trigger("focus") : C.options.liveSearch && C.$searchbox.trigger("focus"), l && (s != I(C.$element[0]) && C.multiple || r != C.$element.prop("selectedIndex") && !C.multiple) && ($ = [d.index, u.prop("selected"), s], C.$element.triggerNative("change"))
					}
				}), this.$menu.on("click", "li." + R.DISABLED + " a, ." + R.POPOVERHEADER + ", ." + R.POPOVERHEADER + " :not(.close)", function(t) {
					t.currentTarget == this && (t.preventDefault(), t.stopPropagation(), C.options.liveSearch && !k(t.target).hasClass("close") ? C.$searchbox.trigger("focus") : C.$button.trigger("focus"))
				}), this.$menuInner.on("click", ".divider, .dropdown-header", function(t) {
					t.preventDefault(), t.stopPropagation(), C.options.liveSearch ? C.$searchbox.trigger("focus") : C.$button.trigger("focus")
				}), this.$menu.on("click", "." + R.POPOVERHEADER + " .close", function() {
					C.$button.trigger("click")
				}), this.$searchbox.on("click", function(t) {
					t.stopPropagation()
				}), this.$menu.on("click", ".actions-btn", function(t) {
					C.options.liveSearch ? C.$searchbox.trigger("focus") : C.$button.trigger("focus"), t.preventDefault(), t.stopPropagation(), k(this).hasClass("bs-select-all") ? C.selectAll() : C.deselectAll()
				}), this.$element.on("change" + j, function() {
					C.render(), C.$element.trigger("changed" + j, $), $ = null
				}).on("focus" + j, function() {
					C.options.mobile || C.$button.trigger("focus")
				})
			},
			liveSearchListener: function() {
				var p = this,
					f = document.createElement("li");
				this.$button.on("click.bs.dropdown.data-api", function() {
					p.$searchbox.val() && p.$searchbox.val("")
				}), this.$searchbox.on("click.bs.dropdown.data-api focus.bs.dropdown.data-api touchend.bs.dropdown.data-api", function(t) {
					t.stopPropagation()
				}), this.$searchbox.on("input propertychange", function() {
					var t = p.$searchbox.val();
					if (p.selectpicker.search.elements = [], p.selectpicker.search.data = [], t) {
						var e = [],
							n = t.toUpperCase(),
							i = {},
							o = [],
							a = p._searchStyle(),
							s = p.options.liveSearchNormalize;
						s && (n = y(n)), p._$lisSelected = p.$menuInner.find(".selected");
						for (var r = 0; r < p.selectpicker.main.data.length; r++) {
							var l = p.selectpicker.main.data[r];
							i[r] || (i[r] = S(l, n, a, s)), i[r] && void 0 !== l.headerIndex && -1 === o.indexOf(l.headerIndex) && (0 < l.headerIndex && (i[l.headerIndex - 1] = !0, o.push(l.headerIndex - 1)), i[l.headerIndex] = !0, o.push(l.headerIndex), i[l.lastIndex + 1] = !0), i[r] && "optgroup-label" !== l.type && o.push(r)
						}
						r = 0;
						for (var c = o.length; r < c; r++) {
							var d = o[r],
								u = o[r - 1],
								h = (l = p.selectpicker.main.data[d], p.selectpicker.main.data[u]);
							("divider" !== l.type || "divider" === l.type && h && "divider" !== h.type && c - 1 !== r) && (p.selectpicker.search.data.push(l), e.push(p.selectpicker.main.elements[d]))
						}
						p.activeIndex = void 0, p.noScroll = !0, p.$menuInner.scrollTop(0), p.selectpicker.search.elements = e, p.createView(!0), e.length || (f.className = "no-results", f.innerHTML = p.options.noneResultsText.replace("{0}", '"' + T(t) + '"'), p.$menuInner[0].firstChild.appendChild(f))
					} else p.$menuInner.scrollTop(0), p.createView(!1)
				})
			},
			_searchStyle: function() {
				return this.options.liveSearchStyle || "contains"
			},
			val: function(t) {
				if (void 0 === t) return this.$element.val();
				var e = I(this.$element[0]);
				return $ = [null, null, e], this.$element.val(t).trigger("changed" + j, $), this.render(), $ = null, this.$element
			},
			changeAll: function(t) {
				if (this.multiple) {
					void 0 === t && (t = !0);
					var e = this.$element[0],
						n = 0,
						i = 0,
						o = I(e);
					e.classList.add("bs-select-hidden");
					for (var a = 0, s = this.selectpicker.current.elements.length; a < s; a++) {
						var r = this.selectpicker.current.data[a],
							l = r.option;
						l && !r.disabled && "divider" !== r.type && (r.selected && n++, (l.selected = t) && i++)
					}
					e.classList.remove("bs-select-hidden"), n !== i && (this.setOptionStatus(), this.togglePlaceholder(), $ = [null, null, o], this.$element.triggerNative("change"))
				}
			},
			selectAll: function() {
				return this.changeAll(!0)
			},
			deselectAll: function() {
				return this.changeAll(!1)
			},
			toggle: function(t) {
				(t = t || window.event) && t.stopPropagation(), this.$button.trigger("click.bs.dropdown.data-api")
			},
			keydown: function(t) {
				var e, n, i, o, a, s = k(this),
					r = s.hasClass("dropdown-toggle"),
					l = (r ? s.closest(".dropdown") : s.closest(U.MENU)).data("this"),
					c = l.findLis(),
					d = !1,
					u = t.which === P && !r && !l.options.selectOnTab,
					h = Y.test(t.which) || u,
					p = l.$menuInner[0].scrollTop,
					f = l.isVirtual(),
					m = !0 === f ? l.selectpicker.view.position0 : 0;
				if (!(n = l.$newElement.hasClass(R.SHOW)) && (h || 48 <= t.which && t.which <= 57 || 96 <= t.which && t.which <= 105 || 65 <= t.which && t.which <= 90) && (l.$button.trigger("click.bs.dropdown.data-api"), l.options.liveSearch)) l.$searchbox.trigger("focus");
				else {
					if (t.which === N && n && (t.preventDefault(), l.$button.trigger("click.bs.dropdown.data-api").trigger("focus")), h) {
						if (!c.length) return;
						void 0 === (e = !0 === f ? c.index(c.filter(".active")) : l.activeIndex) && (e = -1), -1 !== e && ((i = l.selectpicker.current.elements[e + m]).classList.remove("active"), i.firstChild && i.firstChild.classList.remove("active")), t.which === H ? (-1 !== e && e--, e + m < 0 && (e += c.length), l.selectpicker.view.canHighlight[e + m] || -1 === (e = l.selectpicker.view.canHighlight.slice(0, e + m).lastIndexOf(!0) - m) && (e = c.length - 1)) : t.which !== F && !u || (++e + m >= l.selectpicker.view.canHighlight.length && (e = 0), l.selectpicker.view.canHighlight[e + m] || (e = e + 1 + l.selectpicker.view.canHighlight.slice(e + m + 1).indexOf(!0))), t.preventDefault();
						var g = m + e;
						t.which === H ? 0 === m && e === c.length - 1 ? (l.$menuInner[0].scrollTop = l.$menuInner[0].scrollHeight, g = l.selectpicker.current.elements.length - 1) : d = (a = (o = l.selectpicker.current.data[g]).position - o.height) < p : t.which !== F && !u || (0 === e ? g = l.$menuInner[0].scrollTop = 0 : d = p < (a = (o = l.selectpicker.current.data[g]).position - l.sizeInfo.menuInnerHeight)), (i = l.selectpicker.current.elements[g]) && (i.classList.add("active"), i.firstChild && i.firstChild.classList.add("active")), l.activeIndex = l.selectpicker.current.data[g].index, l.selectpicker.view.currentActive = i, d && (l.$menuInner[0].scrollTop = a), l.options.liveSearch ? l.$searchbox.trigger("focus") : s.trigger("focus")
					} else if (!s.is("input") && !X.test(t.which) || t.which === L && l.selectpicker.keydown.keyHistory) {
						var v, b, y = [];
						t.preventDefault(), l.selectpicker.keydown.keyHistory += M[t.which], l.selectpicker.keydown.resetKeyHistory.cancel && clearTimeout(l.selectpicker.keydown.resetKeyHistory.cancel), l.selectpicker.keydown.resetKeyHistory.cancel = l.selectpicker.keydown.resetKeyHistory.start(), b = l.selectpicker.keydown.keyHistory, /^(.)\1+$/.test(b) && (b = b.charAt(0));
						for (var w = 0; w < l.selectpicker.current.data.length; w++) {
							var x = l.selectpicker.current.data[w];
							S(x, b, "startsWith", !0) && l.selectpicker.view.canHighlight[w] && y.push(x.index)
						}
						if (y.length) {
							var E = 0;
							c.removeClass("active").find("a").removeClass("active"), 1 === b.length && (-1 === (E = y.indexOf(l.activeIndex)) || E === y.length - 1 ? E = 0 : E++), v = y[E], d = 0 < p - (o = l.selectpicker.main.data[v]).position ? (a = o.position - o.height, !0) : (a = o.position - l.sizeInfo.menuInnerHeight, o.position > p + l.sizeInfo.menuInnerHeight), (i = l.selectpicker.main.elements[v]).classList.add("active"), i.firstChild && i.firstChild.classList.add("active"), l.activeIndex = y[E], i.firstChild.focus(), d && (l.$menuInner[0].scrollTop = a), s.trigger("focus")
						}
					}
					n && (t.which === L && !l.selectpicker.keydown.keyHistory || t.which === _ || t.which === P && l.options.selectOnTab) && (t.which !== L && t.preventDefault(), l.options.liveSearch && t.which === L || (l.$menuInner.find(".active a").trigger("click", !0), s.trigger("focus"), l.options.liveSearch || (t.preventDefault(), k(document).data("spaceSelect", !0))))
				}
			},
			mobile: function() {
				this.$element[0].classList.add("mobile-device")
			},
			refresh: function() {
				var t = k.extend({}, this.options, this.$element.data());
				this.options = t, this.checkDisabled(), this.setStyle(), this.render(), this.createLi(), this.setWidth(), this.setSize(!0), this.$element.trigger("refreshed" + j)
			},
			hide: function() {
				this.$newElement.hide()
			},
			show: function() {
				this.$newElement.show()
			},
			remove: function() {
				this.$newElement.remove(), this.$element.remove()
			},
			destroy: function() {
				this.$newElement.before(this.$element).remove(), this.$bsContainer ? this.$bsContainer.remove() : this.$menu.remove(), this.$element.off(j).removeData("selectpicker").removeClass("bs-select-hidden selectpicker"), k(window).off(j + "." + this.selectId)
			}
		};
		var tt = k.fn.selectpicker;
		k.fn.selectpicker = Q, k.fn.selectpicker.Constructor = Z, k.fn.selectpicker.noConflict = function() {
			return k.fn.selectpicker = tt, this
		}, k(document).off("keydown.bs.dropdown.data-api").on("keydown" + j, '.bootstrap-select [data-toggle="dropdown"], .bootstrap-select [role="listbox"], .bootstrap-select .bs-searchbox input', Z.prototype.keydown).on("focusin.modal", '.bootstrap-select [data-toggle="dropdown"], .bootstrap-select [role="listbox"], .bootstrap-select .bs-searchbox input', function(t) {
			t.stopPropagation()
		}), k(window).on("load" + j + ".data-api", function() {
			k(".selectpicker").each(function() {
				var t = k(this);
				Q.call(t, t.data())
			})
		})
	}(t)
}),
function(t, e) {
	if ("function" == typeof define && define.amd) define(["module", "exports"], e);
	else if ("undefined" != typeof exports) e(module, exports);
	else {
		var n = {
			exports: {}
		};
		e(n, n.exports), t.autosize = n.exports
	}
}(this, function(t, e) {
	"use strict";
	var n, i, u = "function" == typeof Map ? new Map : (n = [], i = [], {
			has: function(t) {
				return -1 < n.indexOf(t)
			},
			get: function(t) {
				return i[n.indexOf(t)]
			},
			set: function(t, e) {
				-1 === n.indexOf(t) && (n.push(t), i.push(e))
			},
			delete: function(t) {
				var e = n.indexOf(t); - 1 < e && (n.splice(e, 1), i.splice(e, 1))
			}
		}),
		h = function(t) {
			return new Event(t, {
				bubbles: !0
			})
		};
	try {
		new Event("test")
	} catch (t) {
		h = function(t) {
			var e = document.createEvent("Event");
			return e.initEvent(t, !0, !1), e
		}
	}

	function o(o) {
		if (o && o.nodeName && "TEXTAREA" === o.nodeName && !u.has(o)) {
			var t, n = null,
				i = null,
				a = null,
				s = function() {
					o.clientWidth !== i && d()
				},
				r = function(e) {
					window.removeEventListener("resize", s, !1), o.removeEventListener("input", d, !1), o.removeEventListener("keyup", d, !1), o.removeEventListener("autosize:destroy", r, !1), o.removeEventListener("autosize:update", d, !1), Object.keys(e).forEach(function(t) {
						o.style[t] = e[t]
					}), u.delete(o)
				}.bind(o, {
					height: o.style.height,
					resize: o.style.resize,
					overflowY: o.style.overflowY,
					overflowX: o.style.overflowX,
					wordWrap: o.style.wordWrap
				});
			o.addEventListener("autosize:destroy", r, !1), "onpropertychange" in o && "oninput" in o && o.addEventListener("keyup", d, !1), window.addEventListener("resize", s, !1), o.addEventListener("input", d, !1), o.addEventListener("autosize:update", d, !1), o.style.overflowX = "hidden", o.style.wordWrap = "break-word", u.set(o, {
				destroy: r,
				update: d
			}), "vertical" === (t = window.getComputedStyle(o, null)).resize ? o.style.resize = "none" : "both" === t.resize && (o.style.resize = "horizontal"), n = "content-box" === t.boxSizing ? -(parseFloat(t.paddingTop) + parseFloat(t.paddingBottom)) : parseFloat(t.borderTopWidth) + parseFloat(t.borderBottomWidth), isNaN(n) && (n = 0), d()
		}

		function l(t) {
			var e = o.style.width;
			o.style.width = "0px", o.offsetWidth, o.style.width = e, o.style.overflowY = t
		}

		function c() {
			if (0 !== o.scrollHeight) {
				var t = function(t) {
						for (var e = []; t && t.parentNode && t.parentNode instanceof Element;) t.parentNode.scrollTop && e.push({
							node: t.parentNode,
							scrollTop: t.parentNode.scrollTop
						}), t = t.parentNode;
						return e
					}(o),
					e = document.documentElement && document.documentElement.scrollTop;
				o.style.height = "", o.style.height = o.scrollHeight + n + "px", i = o.clientWidth, t.forEach(function(t) {
					t.node.scrollTop = t.scrollTop
				}), e && (document.documentElement.scrollTop = e)
			}
		}

		function d() {
			c();
			var t = Math.round(parseFloat(o.style.height)),
				e = window.getComputedStyle(o, null),
				n = "content-box" === e.boxSizing ? Math.round(parseFloat(e.height)) : o.offsetHeight;
			if (n < t ? "hidden" === e.overflowY && (l("scroll"), c(), n = "content-box" === e.boxSizing ? Math.round(parseFloat(window.getComputedStyle(o, null).height)) : o.offsetHeight) : "hidden" !== e.overflowY && (l("hidden"), c(), n = "content-box" === e.boxSizing ? Math.round(parseFloat(window.getComputedStyle(o, null).height)) : o.offsetHeight), a !== n) {
				a = n;
				var i = h("autosize:resized");
				try {
					o.dispatchEvent(i)
				} catch (t) {}
			}
		}
	}

	function a(t) {
		var e = u.get(t);
		e && e.destroy()
	}

	function s(t) {
		var e = u.get(t);
		e && e.update()
	}
	var r = null;
	"undefined" == typeof window || "function" != typeof window.getComputedStyle ? ((r = function(t) {
		return t
	}).destroy = function(t) {
		return t
	}, r.update = function(t) {
		return t
	}) : ((r = function(t) {
		return t && Array.prototype.forEach.call(t.length ? t : [t], function(t) {
			return o(t)
		}), t
	}).destroy = function(t) {
		return t && Array.prototype.forEach.call(t.length ? t : [t], a), t
	}, r.update = function(t) {
		return t && Array.prototype.forEach.call(t.length ? t : [t], s), t
	}), e.default = r, t.exports = e.default
}),
function(t, e) {
	"object" == typeof exports && "undefined" != typeof module ? module.exports = e() : "function" == typeof define && define.amd ? define(e) : (t = t || self).Sortable = e()
}(this, function() {
	"use strict";

	function o(t) {
		return (o = "function" == typeof Symbol && "symbol" == typeof Symbol.iterator ? function(t) {
			return typeof t
		} : function(t) {
			return t && "function" == typeof Symbol && t.constructor === Symbol && t !== Symbol.prototype ? "symbol" : typeof t
		})(t)
	}

	function s() {
		return (s = Object.assign || function(t) {
			for (var e = 1; e < arguments.length; e++) {
				var n = arguments[e];
				for (var i in n) Object.prototype.hasOwnProperty.call(n, i) && (t[i] = n[i])
			}
			return t
		}).apply(this, arguments)
	}

	function V(o) {
		for (var t = 1; t < arguments.length; t++) {
			var a = null != arguments[t] ? arguments[t] : {},
				e = Object.keys(a);
			"function" == typeof Object.getOwnPropertySymbols && (e = e.concat(Object.getOwnPropertySymbols(a).filter(function(t) {
				return Object.getOwnPropertyDescriptor(a, t).enumerable
			}))), e.forEach(function(t) {
				var e, n, i;
				e = o, i = a[n = t], n in e ? Object.defineProperty(e, n, {
					value: i,
					enumerable: !0,
					configurable: !0,
					writable: !0
				}) : e[n] = i
			})
		}
		return o
	}

	function r(t, e) {
		if (null == t) return {};
		var n, i, o = function(t, e) {
			if (null == t) return {};
			var n, i, o = {},
				a = Object.keys(t);
			for (i = 0; i < a.length; i++) n = a[i], 0 <= e.indexOf(n) || (o[n] = t[n]);
			return o
		}(t, e);
		if (Object.getOwnPropertySymbols) {
			var a = Object.getOwnPropertySymbols(t);
			for (i = 0; i < a.length; i++) n = a[i], 0 <= e.indexOf(n) || Object.prototype.propertyIsEnumerable.call(t, n) && (o[n] = t[n])
		}
		return o
	}

	function e(t) {
		return function(t) {
			if (Array.isArray(t)) {
				for (var e = 0, n = new Array(t.length); e < t.length; e++) n[e] = t[e];
				return n
			}
		}(t) || function(t) {
			if (Symbol.iterator in Object(t) || "[object Arguments]" === Object.prototype.toString.call(t)) return Array.from(t)
		}(t) || function() {
			throw new TypeError("Invalid attempt to spread non-iterable instance")
		}()
	}

	function t(t) {
		return !!navigator.userAgent.match(t)
	}
	var w = t(/(?:Trident.*rv[ :]?11\.|msie|iemobile|Windows Phone)/i),
		x = t(/Edge/i),
		l = t(/firefox/i),
		c = t(/safari/i) && !t(/chrome/i) && !t(/android/i),
		n = t(/iP(ad|od|hone)/i),
		a = t(/chrome/i) && t(/android/i),
		d = {
			capture: !1,
			passive: !1
		};

	function u(t, e, n) {
		t.addEventListener(e, n, !w && d)
	}

	function h(t, e, n) {
		t.removeEventListener(e, n, !w && d)
	}

	function p(t, e) {
		if (e && (">" === e[0] && (e = e.substring(1)), t)) try {
			if (t.matches) return t.matches(e);
			if (t.msMatchesSelector) return t.msMatchesSelector(e);
			if (t.webkitMatchesSelector) return t.webkitMatchesSelector(e)
		} catch (t) {
			return
		}
	}

	function Y(t, e, n, i) {
		if (t) {
			n = n || document;
			do {
				if (null != e && (">" === e[0] ? t.parentNode === n && p(t, e) : p(t, e)) || i && t === n) return t;
				if (t === n) break
			} while (t = (o = t).host && o !== document && o.host.nodeType ? o.host : o.parentNode)
		}
		var o;
		return null
	}
	var f, m = /\s+/g;

	function X(t, e, n) {
		if (t && e)
			if (t.classList) t.classList[n ? "add" : "remove"](e);
			else {
				var i = (" " + t.className + " ").replace(m, " ").replace(" " + e + " ", " ");
				t.className = (i + (n ? " " + e : "")).replace(m, " ")
			}
	}

	function K(t, e, n) {
		var i = t && t.style;
		if (i) {
			if (void 0 === n) return document.defaultView && document.defaultView.getComputedStyle ? n = document.defaultView.getComputedStyle(t, "") : t.currentStyle && (n = t.currentStyle), void 0 === e ? n : n[e];
			e in i || -1 !== e.indexOf("webkit") || (e = "-webkit-" + e), i[e] = n + ("string" == typeof n ? "" : "px")
		}
	}

	function g(t, e) {
		var n = "";
		do {
			var i = K(t, "transform");
			i && "none" !== i && (n = i + " " + n)
		} while (!e && (t = t.parentNode));
		var o = window.DOMMatrix || window.WebKitCSSMatrix || window.CSSMatrix;
		return o && new o(n)
	}

	function v(t, e, n) {
		if (t) {
			var i = t.getElementsByTagName(e),
				o = 0,
				a = i.length;
			if (n)
				for (; o < a; o++) n(i[o], o);
			return i
		}
		return []
	}

	function k() {
		return w ? document.documentElement : document.scrollingElement
	}

	function G(t, e, n, i, o) {
		if (t.getBoundingClientRect || t === window) {
			var a, s, r, l, c, d, u;
			if (u = t !== window && t !== k() ? (s = (a = t.getBoundingClientRect()).top, r = a.left, l = a.bottom, c = a.right, d = a.height, a.width) : (r = s = 0, l = window.innerHeight, c = window.innerWidth, d = window.innerHeight, window.innerWidth), (e || n) && t !== window && (o = o || t.parentNode, !w))
				do {
					if (o && o.getBoundingClientRect && ("none" !== K(o, "transform") || n && "static" !== K(o, "position"))) {
						var h = o.getBoundingClientRect();
						s -= h.top + parseInt(K(o, "border-top-width")), r -= h.left + parseInt(K(o, "border-left-width")), l = s + a.height, c = r + a.width;
						break
					}
				} while (o = o.parentNode);
			if (i && t !== window) {
				var p = g(o || t),
					f = p && p.a,
					m = p && p.d;
				p && (l = (s /= m) + (d /= m), c = (r /= f) + (u /= f))
			}
			return {
				top: s,
				left: r,
				bottom: l,
				right: c,
				width: u,
				height: d
			}
		}
	}

	function q(t, e, n, i) {
		for (var o = A(t, !0), a = (e || G(t))[n]; o;) {
			var s = G(o)[i];
			if (!("top" === i || "left" === i ? s <= a : a <= s)) return o;
			if (o === k()) break;
			o = A(o, !1)
		}
		return !1
	}

	function b(t, e, n) {
		for (var i = 0, o = 0, a = t.children; o < a.length;) {
			if ("none" !== a[o].style.display && a[o] !== Ot.ghost && a[o] !== Ot.dragged && Y(a[o], n.draggable, t, !1)) {
				if (i === e) return a[o];
				i++
			}
			o++
		}
		return null
	}

	function J(t, e) {
		for (var n = t.lastElementChild; n && (n === Ot.ghost || "none" === K(n, "display") || e && !p(n, e));) n = n.previousElementSibling;
		return n || null
	}

	function Z(t, e) {
		var n = 0;
		if (!t || !t.parentNode) return -1;
		for (; t = t.previousElementSibling;) "TEMPLATE" === t.nodeName.toUpperCase() || t === Ot.clone || e && !p(t, e) || n++;
		return n
	}

	function y(t) {
		var e = 0,
			n = 0,
			i = k();
		if (t)
			do {
				var o = g(t),
					a = o.a,
					s = o.d;
				e += t.scrollLeft * a, n += t.scrollTop * s
			} while (t !== i && (t = t.parentNode));
		return [e, n]
	}

	function A(t, e) {
		if (!t || !t.getBoundingClientRect) return k();
		var n = t,
			i = !1;
		do {
			if (n.clientWidth < n.scrollWidth || n.clientHeight < n.scrollHeight) {
				var o = K(n);
				if (n.clientWidth < n.scrollWidth && ("auto" == o.overflowX || "scroll" == o.overflowX) || n.clientHeight < n.scrollHeight && ("auto" == o.overflowY || "scroll" == o.overflowY)) {
					if (!n.getBoundingClientRect || n === document.body) return k();
					if (i || e) return n;
					i = !0
				}
			}
		} while (n = n.parentNode);
		return k()
	}

	function E(t, e) {
		return Math.round(t.top) === Math.round(e.top) && Math.round(t.left) === Math.round(e.left) && Math.round(t.height) === Math.round(e.height) && Math.round(t.width) === Math.round(e.width)
	}

	function S(e, n) {
		return function() {
			if (!f) {
				var t = arguments;
				1 === t.length ? e.call(this, t[0]) : e.apply(this, t), f = setTimeout(function() {
					f = void 0
				}, n)
			}
		}
	}

	function Q(t, e, n) {
		t.scrollLeft += e, t.scrollTop += n
	}

	function D(t) {
		var e = window.Polymer,
			n = window.jQuery || window.Zepto;
		return e && e.dom ? e.dom(t).cloneNode(!0) : n ? n(t).clone(!0)[0] : t.cloneNode(!0)
	}

	function C(t, e) {
		K(t, "position", "absolute"), K(t, "top", e.top), K(t, "left", e.left), K(t, "width", e.width), K(t, "height", e.height)
	}

	function I(t) {
		K(t, "position", ""), K(t, "top", ""), K(t, "left", ""), K(t, "width", ""), K(t, "height", "")
	}
	var tt = "Sortable" + (new Date).getTime();

	function $() {
		var e, i = [];
		return {
			captureAnimationState: function() {
				i = [], this.options.animation && [].slice.call(this.el.children).forEach(function(t) {
					if ("none" !== K(t, "display") && t !== Ot.ghost) {
						i.push({
							target: t,
							rect: G(t)
						});
						var e = G(t);
						if (t.thisAnimationDuration) {
							var n = g(t, !0);
							n && (e.top -= n.f, e.left -= n.e)
						}
						t.fromRect = e
					}
				})
			},
			addAnimationState: function(t) {
				i.push(t)
			},
			removeAnimationState: function(t) {
				i.splice(function(t, e) {
					for (var n in t)
						if (t.hasOwnProperty(n))
							for (var i in e)
								if (e.hasOwnProperty(i) && e[i] === t[n][i]) return Number(n);
					return -1
				}(i, {
					target: t
				}), 1)
			},
			animateAll: function(t) {
				var p = this;
				if (!this.options.animation) return clearTimeout(e), void("function" == typeof t && t());
				var f = !1,
					m = 0;
				i.forEach(function(t) {
					var e, n, i, o, a = 0,
						s = t.target,
						r = s.fromRect,
						l = G(s),
						c = s.prevFromRect,
						d = s.prevToRect,
						u = t.rect,
						h = g(s, !0);
					h && (l.top -= h.f, l.left -= h.e), s.toRect = l, (q(s, l, "bottom", "top") || q(s, l, "top", "bottom") || q(s, l, "right", "left") || q(s, l, "left", "right")) && (q(s, u, "bottom", "top") || q(s, u, "top", "bottom") || q(s, u, "right", "left") || q(s, u, "left", "right")) && (q(s, r, "bottom", "top") || q(s, r, "top", "bottom") || q(s, r, "right", "left") || q(s, r, "left", "right")) || (s.thisAnimationDuration && E(c, l) && !E(r, l) && (u.top - l.top) / (u.left - l.left) == (r.top - l.top) / (r.left - l.left) && (e = u, n = c, i = d, o = p.options, a = Math.sqrt(Math.pow(n.top - e.top, 2) + Math.pow(n.left - e.left, 2)) / Math.sqrt(Math.pow(n.top - i.top, 2) + Math.pow(n.left - i.left, 2)) * o.animation), E(l, r) || (s.prevFromRect = r, s.prevToRect = l, a = a || p.options.animation, p.animate(s, u, a)), a && (f = !0, m = Math.max(m, a), clearTimeout(s.animationResetTimer), s.animationResetTimer = setTimeout(function() {
						s.animationTime = 0, s.prevFromRect = null, s.fromRect = null, s.prevToRect = null, s.thisAnimationDuration = null
					}, a), s.thisAnimationDuration = a))
				}), clearTimeout(e), f ? e = setTimeout(function() {
					"function" == typeof t && t()
				}, m) : "function" == typeof t && t(), i = []
			},
			animate: function(t, e, n) {
				if (n) {
					K(t, "transition", ""), K(t, "transform", "");
					var i = G(t),
						o = g(this.el),
						a = o && o.a,
						s = o && o.d,
						r = (e.left - i.left) / (a || 1),
						l = (e.top - i.top) / (s || 1);
					t.animatingX = !!r, t.animatingY = !!l, K(t, "transform", "translate3d(" + r + "px," + l + "px,0)"), t.offsetWidth, K(t, "transition", "transform " + n + "ms" + (this.options.easing ? " " + this.options.easing : "")), K(t, "transform", "translate3d(0,0,0)"), "number" == typeof t.animated && clearTimeout(t.animated), t.animated = setTimeout(function() {
						K(t, "transition", ""), K(t, "transform", ""), t.animated = !1, t.animatingX = !1, t.animatingY = !1
					}, n)
				}
			}
		}
	}
	var T = [],
		O = {
			initializeByDefault: !0
		},
		M = {
			mount: function(t) {
				for (var e in O) !O.hasOwnProperty(e) || e in t || (t[e] = O[e]);
				T.push(t)
			},
			pluginEvent: function(e, n, i) {
				var o = this;
				this.eventCanceled = !1;
				var a = e + "Global";
				T.forEach(function(t) {
					n[t.pluginName] && (n[t.pluginName][a] && (o.eventCanceled = !!n[t.pluginName][a](V({
						sortable: n
					}, i))), n.options[t.pluginName] && n[t.pluginName][e] && (o.eventCanceled = o.eventCanceled || !!n[t.pluginName][e](V({
						sortable: n
					}, i))))
				})
			},
			initializePlugins: function(i, o, a) {
				for (var t in T.forEach(function(t) {
						var e = t.pluginName;
						if (i.options[e] || t.initializeByDefault) {
							var n = new t(i, o);
							(n.sortable = i)[e] = n, s(a, n.options)
						}
					}), i.options)
					if (i.options.hasOwnProperty(t)) {
						var e = this.modifyOption(i, t, i.options[t]);
						void 0 !== e && (i.options[t] = e)
					}
			},
			getEventOptions: function(e, n) {
				var i = {};
				return T.forEach(function(t) {
					"function" == typeof t.eventOptions && s(i, t.eventOptions.call(n, e))
				}), i
			},
			modifyOption: function(e, n, i) {
				var o;
				return T.forEach(function(t) {
					e[t.pluginName] && t.optionListeners && "function" == typeof t.optionListeners[n] && (o = t.optionListeners[n].call(e[t.pluginName], i))
				}), o
			}
		};

	function N(t) {
		var e, n = t.sortable,
			i = t.rootEl,
			o = t.name,
			a = t.targetEl,
			s = t.cloneEl,
			r = t.toEl,
			l = t.fromEl,
			c = t.oldIndex,
			d = t.newIndex,
			u = t.oldDraggableIndex,
			h = t.newDraggableIndex,
			p = t.originalEvent,
			f = t.putSortable,
			m = t.eventOptions,
			g = (n = n || i[tt]).options,
			v = "on" + o.charAt(0).toUpperCase() + o.substr(1);
		!window.CustomEvent || w || x ? (e = document.createEvent("Event")).initEvent(o, !0, !0) : e = new CustomEvent(o, {
			bubbles: !0,
			cancelable: !0
		}), e.to = r || i, e.from = l || i, e.item = a || i, e.clone = s, e.oldIndex = c, e.newIndex = d, e.oldDraggableIndex = u, e.newDraggableIndex = h, e.originalEvent = p, e.pullMode = f ? f.lastPutMode : void 0;
		var b = V({}, m, M.getEventOptions(o, n));
		for (var y in b) e[y] = b[y];
		i && i.dispatchEvent(e), g[v] && g[v].call(n, e)
	}

	function et(t, e, n) {
		var i = 2 < arguments.length && void 0 !== n ? n : {},
			o = i.evt,
			a = r(i, ["evt"]);
		M.pluginEvent.bind(Ot)(t, e, V({
			dragEl: it,
			parentEl: ot,
			ghostEl: at,
			rootEl: st,
			nextEl: rt,
			lastDownEl: _,
			cloneEl: L,
			cloneHidden: z,
			dragStarted: W,
			putSortable: ut,
			activeSortable: Ot.active,
			originalEvent: o,
			oldIndex: P,
			oldDraggableIndex: H,
			newIndex: lt,
			newDraggableIndex: ct,
			hideGhostForTarget: $t,
			unhideGhostForTarget: Tt,
			cloneNowHidden: function() {
				z = !0
			},
			cloneNowShown: function() {
				z = !1
			},
			dispatchSortableEvent: function(t) {
				nt({
					sortable: e,
					name: t,
					originalEvent: o
				})
			}
		}, a))
	}

	function nt(t) {
		N(V({
			putSortable: ut,
			cloneEl: L,
			targetEl: it,
			rootEl: st,
			oldIndex: P,
			oldDraggableIndex: H,
			newIndex: lt,
			newDraggableIndex: ct
		}, t))
	}
	if ("undefined" == typeof window || !window.document) throw new Error("Sortable.js requires a window with a document");
	var it, ot, at, st, rt, _, L, z, P, lt, H, ct, dt, ut, F, B, W, ht, pt, ft, j, R = !1,
		mt = !1,
		U = [],
		gt = !1,
		vt = !1,
		bt = [],
		yt = !1,
		wt = [],
		xt = n,
		Et = x || w ? "cssFloat" : "float",
		St = !a && !n && "draggable" in document.createElement("div"),
		Dt = function() {
			if (w) return !1;
			var t = document.createElement("x");
			return t.style.cssText = "pointer-events:auto", "auto" === t.style.pointerEvents
		}(),
		Ct = function(t, e) {
			var n = K(t),
				i = parseInt(n.width) - parseInt(n.paddingLeft) - parseInt(n.paddingRight) - parseInt(n.borderLeftWidth) - parseInt(n.borderRightWidth),
				o = b(t, 0, e),
				a = b(t, 1, e),
				s = o && K(o),
				r = a && K(a),
				l = s && parseInt(s.marginLeft) + parseInt(s.marginRight) + G(o).width,
				c = r && parseInt(r.marginLeft) + parseInt(r.marginRight) + G(a).width;
			if ("flex" === n.display) return "column" === n.flexDirection || "column-reverse" === n.flexDirection ? "vertical" : "horizontal";
			if ("grid" === n.display) return n.gridTemplateColumns.split(" ").length <= 1 ? "vertical" : "horizontal";
			if (o && "none" !== s.float) {
				var d = "left" === s.float ? "left" : "right";
				return !a || "both" !== r.clear && r.clear !== d ? "horizontal" : "vertical"
			}
			return o && ("block" === s.display || "flex" === s.display || "table" === s.display || "grid" === s.display || i <= l && "none" === n[Et] || a && "none" === n[Et] && i < l + c) ? "vertical" : "horizontal"
		},
		It = function(t) {
			function l(s, r) {
				return function(t, e, n, i) {
					var o = t.options.group.name && e.options.group.name && t.options.group.name === e.options.group.name;
					if (null == s && (r || o)) return !0;
					if (null == s || !1 === s) return !1;
					if (r && "clone" === s) return s;
					if ("function" == typeof s) return l(s(t, e, n, i), r)(t, e, n, i);
					var a = (r ? t : e).options.group.name;
					return !0 === s || "string" == typeof s && s === a || s.join && -1 < s.indexOf(a)
				}
			}
			var e = {},
				n = t.group;
			n && "object" == o(n) || (n = {
				name: n
			}), e.name = n.name, e.checkPull = l(n.pull, !0), e.checkPut = l(n.put), e.revertClone = n.revertClone, t.group = e
		},
		$t = function() {
			!Dt && at && K(at, "display", "none")
		},
		Tt = function() {
			!Dt && at && K(at, "display", "")
		};
	document.addEventListener("click", function(t) {
		if (mt) return t.preventDefault(), t.stopPropagation && t.stopPropagation(), t.stopImmediatePropagation && t.stopImmediatePropagation(), mt = !1
	}, !0);

	function kt(t) {
		if (it) {
			t = t.touches ? t.touches[0] : t;
			var e = (a = t.clientX, s = t.clientY, U.some(function(t) {
				if (!J(t)) {
					var e = G(t),
						n = t[tt].options.emptyInsertThreshold,
						i = a >= e.left - n && a <= e.right + n,
						o = s >= e.top - n && s <= e.bottom + n;
					return n && i && o ? r = t : void 0
				}
			}), r);
			if (e) {
				var n = {};
				for (var i in t) t.hasOwnProperty(i) && (n[i] = t[i]);
				n.target = n.rootEl = e, n.preventDefault = void 0, n.stopPropagation = void 0, e[tt]._onDragOver(n)
			}
		}
		var a, s, r
	}

	function At(t) {
		it && it.parentNode[tt]._isOutsideThisEl(t.target)
	}

	function Ot(t, e) {
		if (!t || !t.nodeType || 1 !== t.nodeType) throw "Sortable: el must be an HTMLElement, not ".concat({}.toString.call(t));
		this.el = t, this.options = e = s({}, e), t[tt] = this;
		var n = {
			group: null,
			sort: !0,
			disabled: !1,
			store: null,
			handle: null,
			draggable: /^[uo]l$/i.test(t.nodeName) ? ">li" : ">*",
			swapThreshold: 1,
			invertSwap: !1,
			invertedSwapThreshold: null,
			removeCloneOnHide: !0,
			direction: function() {
				return Ct(t, this.options)
			},
			ghostClass: "sortable-ghost",
			chosenClass: "sortable-chosen",
			dragClass: "sortable-drag",
			ignore: "a, img",
			filter: null,
			preventOnFilter: !0,
			animation: 0,
			easing: null,
			setData: function(t, e) {
				t.setData("Text", e.textContent)
			},
			dropBubble: !1,
			dragoverBubble: !1,
			dataIdAttr: "data-id",
			delay: 0,
			delayOnTouchOnly: !1,
			touchStartThreshold: (Number.parseInt ? Number : window).parseInt(window.devicePixelRatio, 10) || 1,
			forceFallback: !1,
			fallbackClass: "sortable-fallback",
			fallbackOnBody: !1,
			fallbackTolerance: 0,
			fallbackOffset: {
				x: 0,
				y: 0
			},
			supportPointer: !1 !== Ot.supportPointer && "PointerEvent" in window,
			emptyInsertThreshold: 5
		};
		for (var i in M.initializePlugins(this, t, n), n) i in e || (e[i] = n[i]);
		for (var o in It(e), this) "_" === o.charAt(0) && "function" == typeof this[o] && (this[o] = this[o].bind(this));
		this.nativeDraggable = !e.forceFallback && St, this.nativeDraggable && (this.options.touchStartThreshold = 1), e.supportPointer ? u(t, "pointerdown", this._onTapStart) : (u(t, "mousedown", this._onTapStart), u(t, "touchstart", this._onTapStart)), this.nativeDraggable && (u(t, "dragover", this), u(t, "dragenter", this)), U.push(this.el), e.store && e.store.get && this.sort(e.store.get(this) || []), s(this, $())
	}

	function Mt(t, e, n, i, o, a, s, r) {
		var l, c, d = t[tt],
			u = d.options.onMove;
		return !window.CustomEvent || w || x ? (l = document.createEvent("Event")).initEvent("move", !0, !0) : l = new CustomEvent("move", {
			bubbles: !0,
			cancelable: !0
		}), l.to = e, l.from = t, l.dragged = n, l.draggedRect = i, l.related = o || e, l.relatedRect = a || G(e), l.willInsertAfter = r, l.originalEvent = s, t.dispatchEvent(l), u && (c = u.call(d, l, s)), c
	}

	function Nt(t) {
		t.draggable = !1
	}

	function _t() {
		yt = !1
	}

	function Lt(t) {
		for (var e = t.tagName + t.className + t.src + t.href + t.textContent, n = e.length, i = 0; n--;) i += e.charCodeAt(n);
		return i.toString(36)
	}

	function zt(t) {
		return setTimeout(t, 0)
	}

	function Pt(t) {
		return clearTimeout(t)
	}
	Ot.prototype = {
		constructor: Ot,
		_isOutsideThisEl: function(t) {
			this.el.contains(t) || t === this.el || (ht = null)
		},
		_getDirection: function(t, e) {
			return "function" == typeof this.options.direction ? this.options.direction.call(this, t, e, it) : this.options.direction
		},
		_onTapStart: function(e) {
			if (e.cancelable) {
				var n = this,
					i = this.el,
					t = this.options,
					o = t.preventOnFilter,
					a = e.type,
					s = e.touches && e.touches[0],
					r = (s || e).target,
					l = e.target.shadowRoot && (e.path && e.path[0] || e.composedPath && e.composedPath()[0]) || r,
					c = t.filter;
				if (! function(t) {
						wt.length = 0;
						var e = t.getElementsByTagName("input"),
							n = e.length;
						for (; n--;) {
							var i = e[n];
							i.checked && wt.push(i)
						}
					}(i), !it && !(/mousedown|pointerdown/.test(a) && 0 !== e.button || t.disabled || l.isContentEditable || (r = Y(r, t.draggable, i, !1)) && r.animated || _ === r)) {
					if (P = Z(r), H = Z(r, t.draggable), "function" == typeof c) {
						if (c.call(this, e, r, this)) return nt({
							sortable: n,
							rootEl: l,
							name: "filter",
							targetEl: r,
							toEl: i,
							fromEl: i
						}), et("filter", n, {
							evt: e
						}), void(o && e.cancelable && e.preventDefault())
					} else if (c = c && c.split(",").some(function(t) {
							if (t = Y(l, t.trim(), i, !1)) return nt({
								sortable: n,
								rootEl: t,
								name: "filter",
								targetEl: r,
								fromEl: i,
								toEl: i
							}), et("filter", n, {
								evt: e
							}), !0
						})) return void(o && e.cancelable && e.preventDefault());
					t.handle && !Y(l, t.handle, i, !1) || this._prepareDragStart(e, s, r)
				}
			}
		},
		_prepareDragStart: function(t, e, n) {
			var i, o = this,
				a = o.el,
				s = o.options,
				r = a.ownerDocument;
			if (n && !it && n.parentNode === a)
				if (st = a, ot = (it = n).parentNode, rt = it.nextSibling, _ = n, dt = s.group, F = {
						target: Ot.dragged = it,
						clientX: (e || t).clientX,
						clientY: (e || t).clientY
					}, this._lastX = (e || t).clientX, this._lastY = (e || t).clientY, it.style["will-change"] = "all", i = function() {
						et("delayEnded", o, {
							evt: t
						}), Ot.eventCanceled ? o._onDrop() : (o._disableDelayedDragEvents(), !l && o.nativeDraggable && (it.draggable = !0), o._triggerDragStart(t, e), nt({
							sortable: o,
							name: "choose",
							originalEvent: t
						}), X(it, s.chosenClass, !0))
					}, s.ignore.split(",").forEach(function(t) {
						v(it, t.trim(), Nt)
					}), u(r, "dragover", kt), u(r, "mousemove", kt), u(r, "touchmove", kt), u(r, "mouseup", o._onDrop), u(r, "touchend", o._onDrop), u(r, "touchcancel", o._onDrop), l && this.nativeDraggable && (this.options.touchStartThreshold = 4, it.draggable = !0), et("delayStart", this, {
						evt: t
					}), !s.delay || s.delayOnTouchOnly && !e || this.nativeDraggable && (x || w)) i();
				else {
					if (Ot.eventCanceled) return void this._onDrop();
					u(r, "mouseup", o._disableDelayedDrag), u(r, "touchend", o._disableDelayedDrag), u(r, "touchcancel", o._disableDelayedDrag), u(r, "mousemove", o._delayedDragTouchMoveHandler), u(r, "touchmove", o._delayedDragTouchMoveHandler), s.supportPointer && u(r, "pointermove", o._delayedDragTouchMoveHandler), o._dragStartTimer = setTimeout(i, s.delay)
				}
		},
		_delayedDragTouchMoveHandler: function(t) {
			var e = t.touches ? t.touches[0] : t;
			Math.max(Math.abs(e.clientX - this._lastX), Math.abs(e.clientY - this._lastY)) >= Math.floor(this.options.touchStartThreshold / (this.nativeDraggable && window.devicePixelRatio || 1)) && this._disableDelayedDrag()
		},
		_disableDelayedDrag: function() {
			it && Nt(it), clearTimeout(this._dragStartTimer), this._disableDelayedDragEvents()
		},
		_disableDelayedDragEvents: function() {
			var t = this.el.ownerDocument;
			h(t, "mouseup", this._disableDelayedDrag), h(t, "touchend", this._disableDelayedDrag), h(t, "touchcancel", this._disableDelayedDrag), h(t, "mousemove", this._delayedDragTouchMoveHandler), h(t, "touchmove", this._delayedDragTouchMoveHandler), h(t, "pointermove", this._delayedDragTouchMoveHandler)
		},
		_triggerDragStart: function(t, e) {
			e = e || "touch" == t.pointerType && t, !this.nativeDraggable || e ? this.options.supportPointer ? u(document, "pointermove", this._onTouchMove) : u(document, e ? "touchmove" : "mousemove", this._onTouchMove) : (u(it, "dragend", this), u(st, "dragstart", this._onDragStart));
			try {
				document.selection ? zt(function() {
					document.selection.empty()
				}) : window.getSelection().removeAllRanges()
			} catch (t) {}
		},
		_dragStarted: function(t, e) {
			if (R = !1, st && it) {
				et("dragStarted", this, {
					evt: e
				}), this.nativeDraggable && u(document, "dragover", At);
				var n = this.options;
				t || X(it, n.dragClass, !1), X(it, n.ghostClass, !0), Ot.active = this, t && this._appendGhost(), nt({
					sortable: this,
					name: "start",
					originalEvent: e
				})
			} else this._nulling()
		},
		_emulateDragOver: function() {
			if (B) {
				this._lastX = B.clientX, this._lastY = B.clientY, $t();
				for (var t = document.elementFromPoint(B.clientX, B.clientY), e = t; t && t.shadowRoot && (t = t.shadowRoot.elementFromPoint(B.clientX, B.clientY)) !== e;) e = t;
				if (it.parentNode[tt]._isOutsideThisEl(t), e)
					do {
						if (e[tt]) {
							if (e[tt]._onDragOver({
									clientX: B.clientX,
									clientY: B.clientY,
									target: t,
									rootEl: e
								}) && !this.options.dragoverBubble) break
						}
						t = e
					} while (e = e.parentNode);
				Tt()
			}
		},
		_onTouchMove: function(t) {
			if (F) {
				var e = this.options,
					n = e.fallbackTolerance,
					i = e.fallbackOffset,
					o = t.touches ? t.touches[0] : t,
					a = at && g(at),
					s = at && a && a.a,
					r = at && a && a.d,
					l = xt && j && y(j),
					c = (o.clientX - F.clientX + i.x) / (s || 1) + (l ? l[0] - bt[0] : 0) / (s || 1),
					d = (o.clientY - F.clientY + i.y) / (r || 1) + (l ? l[1] - bt[1] : 0) / (r || 1),
					u = t.touches ? "translate3d(" + c + "px," + d + "px,0)" : "translate(" + c + "px," + d + "px)";
				if (!Ot.active && !R) {
					if (n && Math.max(Math.abs(o.clientX - this._lastX), Math.abs(o.clientY - this._lastY)) < n) return;
					this._onDragStart(t, !0)
				}
				B = o, K(at, "webkitTransform", u), K(at, "mozTransform", u), K(at, "msTransform", u), K(at, "transform", u), t.cancelable && t.preventDefault()
			}
		},
		_appendGhost: function() {
			if (!at) {
				var t = this.options.fallbackOnBody ? document.body : st,
					e = G(it, !0, xt, !0, t),
					n = this.options;
				if (xt) {
					for (j = t;
						"static" === K(j, "position") && "none" === K(j, "transform") && j !== document;) j = j.parentNode;
					j !== document.body && j !== document.documentElement ? (j === document && (j = k()), e.top += j.scrollTop, e.left += j.scrollLeft) : j = k(), bt = y(j)
				}
				X(at = it.cloneNode(!0), n.ghostClass, !1), X(at, n.fallbackClass, !0), X(at, n.dragClass, !0), K(at, "transition", ""), K(at, "transform", ""), K(at, "box-sizing", "border-box"), K(at, "margin", 0), K(at, "top", e.top), K(at, "left", e.left), K(at, "width", e.width), K(at, "height", e.height), K(at, "opacity", "0.8"), K(at, "position", xt ? "absolute" : "fixed"), K(at, "zIndex", "100000"), K(at, "pointerEvents", "none"), Ot.ghost = at, t.appendChild(at)
			}
		},
		_onDragStart: function(t, e) {
			var n = this,
				i = t.dataTransfer,
				o = n.options;
			et("dragStart", this, {
				evt: t
			}), Ot.eventCanceled ? this._onDrop() : (et("setupClone", this), Ot.eventCanceled || ((L = D(it)).draggable = !1, L.style["will-change"] = "", this._hideClone(), X(L, this.options.chosenClass, !1), Ot.clone = L), n.cloneId = zt(function() {
				et("clone", n), Ot.eventCanceled || (n.options.removeCloneOnHide || st.insertBefore(L, it), n._hideClone(), nt({
					sortable: n,
					name: "clone"
				}))
			}), e || X(it, o.dragClass, !0), e ? (mt = !0, n._loopId = setInterval(n._emulateDragOver, 50)) : (h(document, "mouseup", n._onDrop), h(document, "touchend", n._onDrop), h(document, "touchcancel", n._onDrop), i && (i.effectAllowed = "move", o.setData && o.setData.call(n, i, it)), u(document, "drop", n), K(it, "transform", "translateZ(0)")), R = !0, n._dragStartId = zt(n._dragStarted.bind(n, e, t)), u(document, "selectstart", n), W = !0, c && K(document.body, "user-select", "none"))
		},
		_onDragOver: function(n) {
			var i, o, a, s, r = this.el,
				l = n.target,
				e = this.options,
				t = e.group,
				c = Ot.active,
				d = dt === t,
				u = e.sort,
				h = ut || c,
				p = this,
				f = !1;
			if (!yt) {
				if (void 0 !== n.preventDefault && n.cancelable && n.preventDefault(), l = Y(l, e.draggable, r, !0), W("dragOver"), Ot.eventCanceled) return f;
				if (it.contains(n.target) || l.animated && l.animatingX && l.animatingY || p._ignoreWhileAnimating === l) return R(!1);
				if (mt = !1, c && !e.disabled && (d ? u || (a = !st.contains(it)) : ut === this || (this.lastPutMode = dt.checkPull(this, c, it, n)) && t.checkPut(this, c, it, n))) {
					if (s = "vertical" === this._getDirection(n, l), i = G(it), W("dragOverValid"), Ot.eventCanceled) return f;
					if (a) return ot = st, j(), this._hideClone(), W("revert"), Ot.eventCanceled || (rt ? st.insertBefore(it, rt) : st.appendChild(it)), R(!0);
					var m = J(r, e.draggable);
					if (m && (P = n, H = s, B = G(J((F = this).el, F.options.draggable)), !(H ? P.clientX > B.right + 10 || P.clientX <= B.right && P.clientY > B.bottom && P.clientX >= B.left : P.clientX > B.right && P.clientY > B.top || P.clientX <= B.right && P.clientY > B.bottom + 10) || m.animated)) {
						if (l.parentNode === r) {
							o = G(l);
							var g, v, b, y = it.parentNode !== r,
								w = (T = it.animated && it.toRect || i, k = l.animated && l.toRect || o, O = (A = s) ? T.left : T.top, M = A ? T.right : T.bottom, N = A ? T.width : T.height, _ = A ? k.left : k.top, L = A ? k.right : k.bottom, z = A ? k.width : k.height, !(O === _ || M === L || O + N / 2 === _ + z / 2)),
								x = s ? "top" : "left",
								E = q(l, null, "top", "top") || q(it, null, "top", "top"),
								S = E ? E.scrollTop : void 0;
							if (ht !== l && (v = o[x], gt = !1, vt = !w && e.invertSwap || y), 0 !== (g = function(t, e, n, i, o, a, s, r) {
									var l = i ? t.clientY : t.clientX,
										c = i ? n.height : n.width,
										d = i ? n.top : n.left,
										u = i ? n.bottom : n.right,
										h = !1;
									if (!s)
										if (r && ft < c * o) {
											if (!gt && (1 === pt ? d + c * a / 2 < l : l < u - c * a / 2) && (gt = !0), gt) h = !0;
											else if (1 === pt ? l < d + ft : u - ft < l) return -pt
										} else if (d + c * (1 - o) / 2 < l && l < u - c * (1 - o) / 2) return function(t) {
										return Z(it) < Z(t) ? 1 : -1
									}(e);
									if ((h = h || s) && (l < d + c * a / 2 || u - c * a / 2 < l)) return d + c / 2 < l ? 1 : -1;
									return 0
								}(n, l, o, s, w ? 1 : e.swapThreshold, null == e.invertedSwapThreshold ? e.swapThreshold : e.invertedSwapThreshold, vt, ht === l)))
								for (var D = Z(it); D -= g, (b = ot.children[D]) && ("none" === K(b, "display") || b === at););
							if (0 === g || b === l) return R(!1);
							pt = g;
							var C = (ht = l).nextElementSibling,
								I = !1,
								$ = Mt(st, r, it, i, l, o, n, I = 1 === g);
							if (!1 !== $) return 1 !== $ && -1 !== $ || (I = 1 === $), yt = !0, setTimeout(_t, 30), j(), I && !C ? r.appendChild(it) : l.parentNode.insertBefore(it, I ? C : l), E && Q(E, 0, S - E.scrollTop), ot = it.parentNode, void 0 === v || vt || (ft = Math.abs(v - G(l)[x])), U(), R(!0)
						}
					} else {
						if (m === it) return R(!1);
						if (m && r === n.target && (l = m), l && (o = G(l)), !1 !== Mt(st, r, it, i, l, o, n, !!l)) return j(), r.appendChild(it), ot = r, U(), R(!0)
					}
					if (r.contains(it)) return R(!1)
				}
				var T, k, A, O, M, N, _, L, z, P, H, F, B;
				return !1
			}

			function W(t, e) {
				et(t, p, V({
					evt: n,
					isOwner: d,
					axis: s ? "vertical" : "horizontal",
					revert: a,
					dragRect: i,
					targetRect: o,
					canSort: u,
					fromSortable: h,
					target: l,
					completed: R,
					onMove: function(t, e) {
						return Mt(st, r, it, i, t, G(t), n, e)
					},
					changed: U
				}, e))
			}

			function j() {
				W("dragOverAnimationCapture"), p.captureAnimationState(), p !== h && h.captureAnimationState()
			}

			function R(t) {
				return W("dragOverCompleted", {
					insertion: t
				}), t && (d ? c._hideClone() : c._showClone(p), p !== h && (X(it, ut ? ut.options.ghostClass : c.options.ghostClass, !1), X(it, e.ghostClass, !0)), ut !== p && p !== Ot.active ? ut = p : p === Ot.active && ut && (ut = null), h === p && (p._ignoreWhileAnimating = l), p.animateAll(function() {
					W("dragOverAnimationComplete"), p._ignoreWhileAnimating = null
				}), p !== h && (h.animateAll(), h._ignoreWhileAnimating = null)), (l === it && !it.animated || l === r && !l.animated) && (ht = null), e.dragoverBubble || n.rootEl || l === document || (it.parentNode[tt]._isOutsideThisEl(n.target), t || kt(n)), !e.dragoverBubble && n.stopPropagation && n.stopPropagation(), f = !0
			}

			function U() {
				lt = Z(it), ct = Z(it, e.draggable), nt({
					sortable: p,
					name: "change",
					toEl: r,
					newIndex: lt,
					newDraggableIndex: ct,
					originalEvent: n
				})
			}
		},
		_ignoreWhileAnimating: null,
		_offMoveEvents: function() {
			h(document, "mousemove", this._onTouchMove), h(document, "touchmove", this._onTouchMove), h(document, "pointermove", this._onTouchMove), h(document, "dragover", kt), h(document, "mousemove", kt), h(document, "touchmove", kt)
		},
		_offUpEvents: function() {
			var t = this.el.ownerDocument;
			h(t, "mouseup", this._onDrop), h(t, "touchend", this._onDrop), h(t, "pointerup", this._onDrop), h(t, "touchcancel", this._onDrop), h(document, "selectstart", this)
		},
		_onDrop: function(t) {
			var e = this.el,
				n = this.options;
			lt = Z(it), ct = Z(it, n.draggable), et("drop", this, {
				evt: t
			}), lt = Z(it), ct = Z(it, n.draggable), Ot.eventCanceled || (gt = vt = R = !1, clearInterval(this._loopId), clearTimeout(this._dragStartTimer), Pt(this.cloneId), Pt(this._dragStartId), this.nativeDraggable && (h(document, "drop", this), h(e, "dragstart", this._onDragStart)), this._offMoveEvents(), this._offUpEvents(), c && K(document.body, "user-select", ""), t && (W && (t.cancelable && t.preventDefault(), n.dropBubble || t.stopPropagation()), at && at.parentNode && at.parentNode.removeChild(at), (st === ot || ut && "clone" !== ut.lastPutMode) && L && L.parentNode && L.parentNode.removeChild(L), it && (this.nativeDraggable && h(it, "dragend", this), Nt(it), it.style["will-change"] = "", W && !R && X(it, ut ? ut.options.ghostClass : this.options.ghostClass, !1), X(it, this.options.chosenClass, !1), nt({
				sortable: this,
				name: "unchoose",
				toEl: ot,
				newIndex: null,
				newDraggableIndex: null,
				originalEvent: t
			}), st !== ot ? (0 <= lt && (nt({
				rootEl: ot,
				name: "add",
				toEl: ot,
				fromEl: st,
				originalEvent: t
			}), nt({
				sortable: this,
				name: "remove",
				toEl: ot,
				originalEvent: t
			}), nt({
				rootEl: ot,
				name: "sort",
				toEl: ot,
				fromEl: st,
				originalEvent: t
			}), nt({
				sortable: this,
				name: "sort",
				toEl: ot,
				originalEvent: t
			})), ut && ut.save()) : lt !== P && 0 <= lt && (nt({
				sortable: this,
				name: "update",
				toEl: ot,
				originalEvent: t
			}), nt({
				sortable: this,
				name: "sort",
				toEl: ot,
				originalEvent: t
			})), Ot.active && (null != lt && -1 !== lt || (lt = P, ct = H), nt({
				sortable: this,
				name: "end",
				toEl: ot,
				originalEvent: t
			}), this.save())))), this._nulling()
		},
		_nulling: function() {
			et("nulling", this), st = it = ot = at = rt = L = _ = z = F = B = W = lt = ct = P = H = ht = pt = ut = dt = Ot.dragged = Ot.ghost = Ot.clone = Ot.active = null, wt.forEach(function(t) {
				t.checked = !0
			}), wt.length = 0
		},
		handleEvent: function(t) {
			switch (t.type) {
				case "drop":
				case "dragend":
					this._onDrop(t);
					break;
				case "dragenter":
				case "dragover":
					it && (this._onDragOver(t), function(t) {
						t.dataTransfer && (t.dataTransfer.dropEffect = "move");
						t.cancelable && t.preventDefault()
					}(t));
					break;
				case "selectstart":
					t.preventDefault()
			}
		},
		toArray: function() {
			for (var t, e = [], n = this.el.children, i = 0, o = n.length, a = this.options; i < o; i++) Y(t = n[i], a.draggable, this.el, !1) && e.push(t.getAttribute(a.dataIdAttr) || Lt(t));
			return e
		},
		sort: function(t) {
			var i = {},
				o = this.el;
			this.toArray().forEach(function(t, e) {
				var n = o.children[e];
				Y(n, this.options.draggable, o, !1) && (i[t] = n)
			}, this), t.forEach(function(t) {
				i[t] && (o.removeChild(i[t]), o.appendChild(i[t]))
			})
		},
		save: function() {
			var t = this.options.store;
			t && t.set && t.set(this)
		},
		closest: function(t, e) {
			return Y(t, e || this.options.draggable, this.el, !1)
		},
		option: function(t, e) {
			var n = this.options;
			if (void 0 === e) return n[t];
			var i = M.modifyOption(this, t, e);
			n[t] = void 0 !== i ? i : e, "group" === t && It(n)
		},
		destroy: function() {
			et("destroy", this);
			var t = this.el;
			t[tt] = null, h(t, "mousedown", this._onTapStart), h(t, "touchstart", this._onTapStart), h(t, "pointerdown", this._onTapStart), this.nativeDraggable && (h(t, "dragover", this), h(t, "dragenter", this)), Array.prototype.forEach.call(t.querySelectorAll("[draggable]"), function(t) {
				t.removeAttribute("draggable")
			}), this._onDrop(), U.splice(U.indexOf(this.el), 1), this.el = t = null
		},
		_hideClone: function() {
			if (!z) {
				if (et("hideClone", this), Ot.eventCanceled) return;
				K(L, "display", "none"), this.options.removeCloneOnHide && L.parentNode && L.parentNode.removeChild(L), z = !0
			}
		},
		_showClone: function(t) {
			if ("clone" === t.lastPutMode) {
				if (z) {
					if (et("showClone", this), Ot.eventCanceled) return;
					st.contains(it) && !this.options.group.revertClone ? st.insertBefore(L, it) : rt ? st.insertBefore(L, rt) : st.appendChild(L), this.options.group.revertClone && this._animate(it, L), K(L, "display", ""), z = !1
				}
			} else this._hideClone()
		}
	}, u(document, "touchmove", function(t) {
		(Ot.active || R) && t.cancelable && t.preventDefault()
	}), Ot.utils = {
		on: u,
		off: h,
		css: K,
		find: v,
		is: function(t, e) {
			return !!Y(t, e, t, !1)
		},
		extend: function(t, e) {
			if (t && e)
				for (var n in e) e.hasOwnProperty(n) && (t[n] = e[n]);
			return t
		},
		throttle: S,
		closest: Y,
		toggleClass: X,
		clone: D,
		index: Z,
		nextTick: zt,
		cancelNextTick: Pt,
		detectDirection: Ct,
		getChild: b
	}, Ot.mount = function() {
		for (var t = arguments.length, e = new Array(t), n = 0; n < t; n++) e[n] = arguments[n];
		e[0].constructor === Array && (e = e[0]), e.forEach(function(t) {
			if (!t.prototype || !t.prototype.constructor) throw "Sortable: Mounted plugin must be a constructor function, not ".concat({}.toString.call(el));
			t.utils && (Ot.utils = V({}, Ot.utils, t.utils)), M.mount(t)
		})
	}, Ot.create = function(t, e) {
		return new Ot(t, e)
	};
	var Ht, Ft, Bt, Wt, jt, Rt, Ut = [],
		Vt = !(Ot.version = "1.10.0-rc3");

	function Yt() {
		Ut.forEach(function(t) {
			clearInterval(t.pid)
		}), Ut = []
	}

	function Xt() {
		clearInterval(Rt)
	}

	function Kt(t) {
		var e = t.originalEvent,
			n = t.putSortable,
			i = t.dragEl,
			o = t.activeSortable,
			a = t.dispatchSortableEvent,
			s = t.hideGhostForTarget,
			r = t.unhideGhostForTarget,
			l = n || o;
		s();
		var c = document.elementFromPoint(e.clientX, e.clientY);
		r(), l && !l.el.contains(c) && (a("spill"), this.onSpill(i))
	}
	var Gt, qt = S(function(n, t, e, i) {
		if (t.scroll) {
			var o, a = t.scrollSensitivity,
				s = t.scrollSpeed,
				r = k(),
				l = !1;
			Ft !== e && (Ft = e, Yt(), Ht = t.scroll, o = t.scrollFn, !0 === Ht && (Ht = A(e, !0)));
			var c = 0,
				d = Ht;
			do {
				var u = d,
					h = G(u),
					p = h.top,
					f = h.bottom,
					m = h.left,
					g = h.right,
					v = h.width,
					b = h.height,
					y = void 0,
					w = void 0,
					x = u.scrollWidth,
					E = u.scrollHeight,
					S = K(u),
					D = u.scrollLeft,
					C = u.scrollTop;
				w = u === r ? (y = v < x && ("auto" === S.overflowX || "scroll" === S.overflowX || "visible" === S.overflowX), b < E && ("auto" === S.overflowY || "scroll" === S.overflowY || "visible" === S.overflowY)) : (y = v < x && ("auto" === S.overflowX || "scroll" === S.overflowX), b < E && ("auto" === S.overflowY || "scroll" === S.overflowY));
				var I = y && (Math.abs(g - n.clientX) <= a && D + v < x) - (Math.abs(m - n.clientX) <= a && !!D),
					$ = w && (Math.abs(f - n.clientY) <= a && C + b < E) - (Math.abs(p - n.clientY) <= a && !!C);
				if (!Ut[c])
					for (var T = 0; T <= c; T++) Ut[T] || (Ut[T] = {});
				Ut[c].vx == I && Ut[c].vy == $ && Ut[c].el === u || (Ut[c].el = u, Ut[c].vx = I, Ut[c].vy = $, clearInterval(Ut[c].pid), 0 == I && 0 == $ || (l = !0, Ut[c].pid = setInterval(function() {
					i && 0 === this.layer && Ot.active._onTouchMove(jt);
					var t = Ut[this.layer].vy ? Ut[this.layer].vy * s : 0,
						e = Ut[this.layer].vx ? Ut[this.layer].vx * s : 0;
					"function" == typeof o && "continue" !== o.call(Ot.dragged.parentNode[tt], e, t, n, jt, Ut[this.layer].el) || Q(Ut[this.layer].el, e, t)
				}.bind({
					layer: c
				}), 24))), c++
			} while (t.bubbleScroll && d !== r && (d = A(d, !1)));
			Vt = l
		}
	}, 30);

	function Jt() {}

	function Zt() {}
	Jt.prototype = {
		startIndex: null,
		dragStart: function(t) {
			var e = t.oldDraggableIndex;
			this.startIndex = e
		},
		onSpill: function(t) {
			this.sortable.captureAnimationState();
			var e = b(this.sortable.el, this.startIndex, this.sortable.options);
			e ? this.sortable.el.insertBefore(t, e) : this.sortable.el.appendChild(t), this.sortable.animateAll()
		},
		drop: Kt
	}, s(Jt, {
		pluginName: "revertOnSpill"
	}), Zt.prototype = {
		onSpill: function(t) {
			this.sortable.captureAnimationState(), t.parentNode && t.parentNode.removeChild(t), this.sortable.animateAll()
		},
		drop: Kt
	}, s(Zt, {
		pluginName: "removeOnSpill"
	});
	var Qt, te, ee, ne, ie, oe = [],
		ae = [],
		se = !1,
		re = !1,
		le = !1;

	function ce(n, o) {
		ae.forEach(function(t) {
			var e = o.children[t.sortableIndex + (n ? Number(i) : 0)];
			e ? o.insertBefore(t, e) : o.appendChild(t)
		})
	}

	function de() {
		oe.forEach(function(t) {
			t !== ee && t.parentNode && t.parentNode.removeChild(t)
		})
	}
	return Ot.mount(new function() {
		function t() {
			for (var t in this.options = {
					scroll: !0,
					scrollSensitivity: 30,
					scrollSpeed: 10,
					bubbleScroll: !0
				}, this) "_" === t.charAt(0) && "function" == typeof this[t] && (this[t] = this[t].bind(this))
		}
		return t.prototype = {
			dragStarted: function(t) {
				var e = t.originalEvent;
				this.sortable.nativeDraggable ? u(document, "dragover", this._handleAutoScroll) : this.sortable.options.supportPointer ? u(document, "pointermove", this._handleFallbackAutoScroll) : e.touches ? u(document, "touchmove", this._handleFallbackAutoScroll) : u(document, "mousemove", this._handleFallbackAutoScroll)
			},
			dragOverCompleted: function(t) {
				var e = t.originalEvent;
				this.sortable.options.dragOverBubble || e.rootEl || this._handleAutoScroll(e)
			},
			drop: function() {
				this.sortable.nativeDraggable ? h(document, "dragover", this._handleAutoScroll) : (h(document, "pointermove", this._handleFallbackAutoScroll), h(document, "touchmove", this._handleFallbackAutoScroll), h(document, "mousemove", this._handleFallbackAutoScroll)), Xt(), Yt(), clearTimeout(f), f = void 0
			},
			nulling: function() {
				jt = Ft = Ht = Vt = Rt = Bt = Wt = null, Ut.length = 0
			},
			_handleFallbackAutoScroll: function(t) {
				this._handleAutoScroll(t, !0)
			},
			_handleAutoScroll: function(e, n) {
				var i = this,
					o = e.clientX,
					a = e.clientY,
					t = document.elementFromPoint(o, a);
				if (jt = e, n || x || w || c) {
					qt(e, this.options, t, n);
					var s = A(t, !0);
					!Vt || Rt && o === Bt && a === Wt || (Rt && Xt(), Rt = setInterval(function() {
						var t = A(document.elementFromPoint(o, a), !0);
						t !== s && (s = t, Yt()), qt(e, i.options, t, n)
					}, 10), Bt = o, Wt = a)
				} else {
					if (!this.sortable.options.bubbleScroll || A(t, !0) === k()) return void Yt();
					qt(e, this.options, A(t, !1), !1)
				}
			}
		}, s(t, {
			pluginName: "scroll",
			initializeByDefault: !0
		})
	}), Ot.mount(Zt, Jt), Ot.mount(new function() {
		function t() {
			this.options = {
				swapClass: "sortable-swap-highlight"
			}
		}
		return t.prototype = {
			dragStart: function(t) {
				var e = t.dragEl;
				Gt = e
			},
			dragOverValid: function(t) {
				var e = t.completed,
					n = t.target,
					i = t.onMove,
					o = t.activeSortable,
					a = t.changed;
				if (o.options.swap) {
					var s = this.sortable.el,
						r = this.sortable.options;
					if (n && n !== s) {
						var l = Gt;
						Gt = !1 !== i(n) ? (X(n, r.swapClass, !0), n) : null, l && l !== Gt && X(l, r.swapClass, !1)
					}
					return a(), e(!0)
				}
			},
			drop: function(t) {
				var e = t.activeSortable,
					n = t.putSortable,
					i = t.dragEl,
					o = n || this.sortable,
					a = this.sortable.options;
				Gt && X(Gt, a.swapClass, !1), Gt && (a.swap || n && n.options.swap) && i !== Gt && (o.captureAnimationState(), o !== e && e.captureAnimationState(), function(t, e) {
					var n, i, o = t.parentNode,
						a = e.parentNode;
					if (!o || !a || o.isEqualNode(e) || a.isEqualNode(t)) return;
					n = Z(t), i = Z(e), o.isEqualNode(a) && n < i && i++;
					o.insertBefore(e, o.children[n]), a.insertBefore(t, a.children[i])
				}(i, Gt), o.animateAll(), o !== e && e.animateAll())
			},
			nulling: function() {
				Gt = null
			}
		}, s(t, {
			pluginName: "swap",
			eventOptions: function() {
				return {
					swapItem: Gt
				}
			}
		})
	}), Ot.mount(new function() {
		function t(i) {
			for (var t in this) "_" === t.charAt(0) && "function" == typeof this[t] && (this[t] = this[t].bind(this));
			i.options.supportPointer ? u(document, "pointerup", this._deselectMultiDrag) : (u(document, "mouseup", this._deselectMultiDrag), u(document, "touchend", this._deselectMultiDrag)), u(document, "keydown", this._checkKeyDown), u(document, "keyup", this._checkKeyUp), this.options = {
				selectedClass: "sortable-selected",
				multiDragKey: null,
				setData: function(t, e) {
					var n = "";
					oe.length && te === i ? oe.forEach(function(t, e) {
						n += (e ? ", " : "") + t.textContent
					}) : n = e.textContent, t.setData("Text", n)
				}
			}
		}
		return t.prototype = {
			multiDragKeyDown: !1,
			isMultiDrag: !1,
			delayStartGlobal: function(t) {
				var e = t.dragEl;
				ee = e
			},
			delayEnded: function() {
				this.isMultiDrag = ~oe.indexOf(ee)
			},
			setupClone: function(t) {
				var e = t.sortable;
				if (this.isMultiDrag) {
					for (var n = 0; n < oe.length; n++) ae.push(D(oe[n])), ae[n].sortableIndex = oe[n].sortableIndex, ae[n].draggable = !1, ae[n].style["will-change"] = "", X(ae[n], e.options.selectedClass, !1), oe[n] === ee && X(ae[n], e.options.chosenClass, !1);
					return e._hideClone(), !0
				}
			},
			clone: function(t) {
				var e = t.sortable,
					n = t.rootEl,
					i = t.dispatchSortableEvent;
				if (this.isMultiDrag) return !e.options.removeCloneOnHide && oe.length && te === e ? (ce(!0, n), i("clone"), !0) : void 0
			},
			showClone: function(t) {
				var e = t.cloneNowShown,
					n = t.rootEl;
				if (this.isMultiDrag) return ce(!1, n), ae.forEach(function(t) {
					K(t, "display", "")
				}), e(), !(ie = !1)
			},
			hideClone: function(t) {
				var e = t.sortable,
					n = t.cloneNowHidden;
				if (this.isMultiDrag) return ae.forEach(function(t) {
					K(t, "display", "none"), e.options.removeCloneOnHide && t.parentNode && t.parentNode.removeChild(t)
				}), n(), ie = !0
			},
			dragStartGlobal: function(t) {
				t.sortable;
				!this.isMultiDrag && te && te.multiDrag._deselectMultiDrag(), oe.forEach(function(t) {
					t.sortableIndex = Z(t)
				}), oe = oe.sort(function(t, e) {
					return t.sortableIndex - e.sortableIndex
				}), le = !0
			},
			dragStarted: function(t) {
				var e = t.sortable;
				if (this.isMultiDrag) {
					if (e.options.sort && (e.captureAnimationState(), e.options.animation)) {
						oe.forEach(function(t) {
							t !== ee && K(t, "position", "absolute")
						});
						var n = G(ee, !1, !0, !0);
						oe.forEach(function(t) {
							t !== ee && C(t, n)
						}), se = re = !0
					}
					e.animateAll(function() {
						se = re = !1, e.options.animation && oe.forEach(function(t) {
							I(t)
						}), e.options.sort && de()
					})
				}
			},
			dragOver: function(t) {
				var e = t.target,
					n = t.completed;
				if (re && ~oe.indexOf(e)) return n(!1)
			},
			revert: function(t) {
				var n, o, e = t.fromSortable,
					a = t.rootEl,
					s = t.sortable,
					r = t.dragRect;
				1 < oe.length && (oe.forEach(function(t) {
					s.addAnimationState({
						target: t,
						rect: re ? G(t) : r
					}), I(t), t.fromRect = r, e.removeAnimationState(t)
				}), re = !1, n = !s.options.removeCloneOnHide, o = a, oe.forEach(function(t) {
					var e = o.children[t.sortableIndex + (n ? Number(i) : 0)];
					e ? o.insertBefore(t, e) : o.appendChild(t)
				}))
			},
			dragOverCompleted: function(t) {
				var e = t.sortable,
					n = t.isOwner,
					i = t.insertion,
					o = t.activeSortable,
					a = t.parentEl,
					s = t.putSortable,
					r = e.options;
				if (i) {
					if (n && o._hideClone(), se = !1, r.animation && 1 < oe.length && (re || !n && !o.options.sort && !s)) {
						var l = G(ee, !1, !0, !0);
						oe.forEach(function(t) {
							t !== ee && (C(t, l), a.appendChild(t))
						}), re = !0
					}
					if (!n)
						if (re || de(), 1 < oe.length) {
							var c = ie;
							o._showClone(e), o.options.animation && !ie && c && ae.forEach(function(t) {
								o.addAnimationState({
									target: t,
									rect: ne
								}), t.fromRect = ne, t.thisAnimationDuration = null
							})
						} else o._showClone(e)
				}
			},
			dragOverAnimationCapture: function(t) {
				var e = t.dragRect,
					n = t.isOwner,
					i = t.activeSortable;
				if (oe.forEach(function(t) {
						t.thisAnimationDuration = null
					}), i.options.animation && !n && i.multiDrag.isMultiDrag) {
					ne = s({}, e);
					var o = g(ee, !0);
					ne.top -= o.f, ne.left -= o.e
				}
			},
			dragOverAnimationComplete: function() {
				re && (re = !1, de())
			},
			drop: function(t) {
				var e = t.originalEvent,
					n = t.rootEl,
					i = t.parentEl,
					o = t.sortable,
					a = t.dispatchSortableEvent,
					s = t.oldIndex,
					r = t.putSortable,
					l = r || this.sortable;
				if (e) {
					var c = o.options,
						d = i.children;
					if (!le)
						if (c.multiDragKey && !this.multiDragKeyDown && this._deselectMultiDrag(), X(ee, c.selectedClass, !~oe.indexOf(ee)), ~oe.indexOf(ee)) oe.splice(oe.indexOf(ee), 1), Qt = null, N({
							sortable: o,
							rootEl: n,
							name: "deselect",
							targetEl: ee,
							originalEvt: e
						});
						else {
							if (oe.push(ee), N({
									sortable: o,
									rootEl: n,
									name: "select",
									targetEl: ee,
									originalEvt: e
								}), (!c.multiDragKey || this.multiDragKeyDown) && e.shiftKey && Qt && o.el.contains(Qt)) {
								var u, h, p = Z(Qt),
									f = Z(ee);
								if (~p && ~f && p !== f)
									for (u = p < f ? (h = p, f) : (h = f, p + 1); h < u; h++) ~oe.indexOf(d[h]) || (X(d[h], c.selectedClass, !0), oe.push(d[h]), N({
										sortable: o,
										rootEl: n,
										name: "select",
										targetEl: d[h],
										originalEvt: e
									}))
							} else Qt = ee;
							te = l
						} if (le && this.isMultiDrag) {
						if ((i[tt].options.sort || i !== n) && 1 < oe.length) {
							var m = G(ee),
								g = Z(ee, ":not(." + this.options.selectedClass + ")");
							if (!se && c.animation && (ee.thisAnimationDuration = null), l.captureAnimationState(), !se && (c.animation && (ee.fromRect = m, oe.forEach(function(t) {
									if (t.thisAnimationDuration = null, t !== ee) {
										var e = re ? G(t) : m;
										t.fromRect = e, l.addAnimationState({
											target: t,
											rect: e
										})
									}
								})), de(), oe.forEach(function(t) {
									d[g] ? i.insertBefore(t, d[g]) : i.appendChild(t), g++
								}), s === Z(ee))) {
								var v = !1;
								oe.forEach(function(t) {
									t.sortableIndex === Z(t) || (v = !0)
								}), v && a("update")
							}
							oe.forEach(function(t) {
								I(t)
							}), l.animateAll()
						}
						te = l
					}(n === i || r && "clone" !== r.lastPutMode) && ae.forEach(function(t) {
						t.parentNode && t.parentNode.removeChild(t)
					})
				}
			},
			nullingGlobal: function() {
				this.isMultiDrag = le = !1, ae.length = 0
			},
			destroy: function() {
				this._deselectMultiDrag(), h(document, "pointerup", this._deselectMultiDrag), h(document, "mouseup", this._deselectMultiDrag), h(document, "touchend", this._deselectMultiDrag), h(document, "keydown", this._checkKeyDown), h(document, "keyup", this._checkKeyUp)
			},
			_deselectMultiDrag: function(t) {
				if (!le && te === this.sortable && !(t && Y(t.target, this.sortable.options.draggable, this.sortable.el, !1) || t && 0 !== t.button))
					for (; oe.length;) {
						var e = oe[0];
						X(e, this.sortable.options.selectedClass, !1), oe.shift(), N({
							sortable: this.sortable,
							rootEl: this.sortable.el,
							name: "deselect",
							targetEl: e,
							originalEvt: t
						})
					}
			},
			_checkKeyDown: function(t) {
				t.key === this.sortable.options.multiDragKey && (this.multiDragKeyDown = !0)
			},
			_checkKeyUp: function(t) {
				t.key === this.sortable.options.multiDragKey && (this.multiDragKeyDown = !1)
			}
		}, s(t, {
			pluginName: "multiDrag",
			utils: {
				select: function(t) {
					var e = t.parentNode[tt];
					e && e.options.multiDrag && !~oe.indexOf(t) && (te && te !== e && (te.multiDrag._deselectMultiDrag(), te = e), X(t, e.options.selectedClass, !0), oe.push(t))
				},
				deselect: function(t) {
					var e = t.parentNode[tt],
						n = oe.indexOf(t);
					e && e.options.multiDrag && ~n && (X(t, e.options.selectedClass, !1), oe.splice(n, 1))
				}
			},
			eventOptions: function() {
				var n = this,
					i = [],
					o = [];
				return oe.forEach(function(t) {
					var e;
					i.push({
						multiDragElement: t,
						index: t.sortableIndex
					}), e = re && t !== ee ? -1 : re ? Z(t, ":not(." + n.options.selectedClass + ")") : Z(t), o.push({
						multiDragElement: t,
						index: e
					})
				}), {
					items: e(oe),
					clones: [].concat(ae),
					oldIndicies: i,
					newIndicies: o
				}
			},
			optionListeners: {
				multiDragKey: function(t) {
					return "ctrl" === (t = t.toLowerCase()) ? t = "Control" : 1 < t.length && (t = t.charAt(0).toUpperCase() + t.substr(1)), t
				}
			}
		})
	}), Ot
}),
function(window, $) {
	var fave = function(window, $) {
		var FormDataWasChanged = !1;

		function IsDebugMode() {
			return window.debug && !0 === window.debug
		}

		function GetModalAlertTmpl(t, e, n) {
			return '<div class="alert alert-' + (n ? "danger" : "success") + ' alert-dismissible fade show" role="alert"><strong>' + t + "</strong> " + e + '<button type="button" class="close" data-dismiss="alert" aria-label="Close"><span aria-hidden="true">&times;</span></button></div>'
		}

		function ShowSystemMsg(t, e, n) {
			var i = $(".modal.show .sys-messages");
			i.length || (i = $("form.alert-here .sys-messages")), i.length ? i.html(GetModalAlertTmpl(t, e, n)) : ShowSystemMsgModal(t, e, n)
		}

		function ShowSystemMsgModal(t, e, n) {
			$("#sys-modal-system-message-placeholder").html("");
			var i = '<div class="modal fade" id="sys-modal-system-message" tabindex="-1" role="dialog" aria-labelledby="sysModalSystemMessageLabel" aria-hidden="true"> \t\t\t\t<div class="modal-dialog modal-dialog-centered" role="document"> \t\t\t\t\t<div class="modal-content"> \t\t\t\t\t\t\t<div class="modal-header"> \t\t\t\t\t\t\t\t<h5 class="modal-title" id="sysModalSystemMessageLabel">' + t + '</h5> \t\t\t\t\t\t\t\t<button type="button" class="close" data-dismiss="modal" aria-label="Close"> \t\t\t\t\t\t\t\t\t<span aria-hidden="true">&times;</span> \t\t\t\t\t\t\t\t</button> \t\t\t\t\t\t\t</div> \t\t\t\t\t\t\t<div class="modal-body text-left">' + e + '</div> \t\t\t\t\t\t\t<div class="modal-footer"> \t\t\t\t\t\t\t\t<button type="button" class="btn btn-secondary" data-dismiss="modal">Cancel</button> \t\t\t\t\t\t\t</div> \t\t\t\t\t</div> \t\t\t\t</div> \t\t\t</div>';
			$("#sys-modal-system-message-placeholder").html(i), $("#sys-modal-system-message").modal({
				backdrop: "static",
				keyboard: !0,
				show: !1
			}), $("#sys-modal-system-message").on("hidden.bs.modal", function(t) {
				$("#sys-modal-system-message-placeholder").html("")
			}), $("#sys-modal-system-message").modal("show")
		}

		function AjaxEval(data) {
			try {
				eval(data)
			} catch (t) {
				t instanceof SyntaxError && (console.log(data), console.log("Error: JavaScript code eval error", t.message))
			}
		}

		function AjaxDone(t) {
			AjaxEval(t)
		}

		function AjaxFail(t, e, n) {
			"error" === e.toLowerCase() && "not found" === n.toLowerCase() ? AjaxEval(t) : (console.log("Error: data sending error, page will be reloaded", t, e, n), setTimeout(function() {
				window.location.reload(!1)
			}, 1e3))
		}

		function FormToAjax(n) {
			n.submit(function(t) {
				if (n.hasClass("loading")) t.preventDefault();
				else {
					n.addClass("loading").addClass("alert-here");
					var e = n.find("button[type=submit]");
					e.addClass("progress-bar-striped").addClass("progress-bar-animated"), "" != e.attr("data-target") && $("#" + e.attr("data-target")).addClass("progress-bar-striped").addClass("progress-bar-animated"), n.find(".sys-messages").html(""), $.ajax({
						type: "POST",
						url: n.attr("action"),
						data: n.serialize()
					}).done(function(t) {
						FormDataWasChanged = !1, IsDebugMode() && console.log("done", t), AjaxDone(t)
					}).fail(function(t, e, n) {
						IsDebugMode() && console.log("fail", t, e, n), AjaxFail(t.responseText, e, n)
					}).always(function() {
						setTimeout(function() {
							n.removeClass("loading").removeClass("alert-here"), e.removeClass("progress-bar-striped").removeClass("progress-bar-animated"), "" != e.attr("data-target") && $("#" + e.attr("data-target")).removeClass("progress-bar-striped").removeClass("progress-bar-animated")
						}, 100)
					}), t.preventDefault()
				}
			});
			var t = n.find("button[type=submit]");
			"" != t.attr("data-target") && $("#" + t.attr("data-target")).click(function() {
				t.click()
			}), n.hasClass("prev-data-lost") && n.find("input, textarea, select").on("input", function() {
				FormDataWasChanged || $(this).hasClass("ignore-lost-data") || (FormDataWasChanged = !0)
			})
		}

		function PreventDataLost() {
			FormDataWasChanged = FormDataWasChanged || !0
		}

		function FormDataIsChanged() {
			return FormDataWasChanged
		}

		function HtmlDecode(t) {
			return (new DOMParser).parseFromString(t, "text/html").documentElement.textContent
		}

		function HtmlFixEditorHtml(t) {
			return newValue = t, newValue = newValue.replace(/&nbsp;/gi, ""), newValue
		}

		function AllFormsToAjax() {
			$("form").each(function() {
				FormToAjax($(this))
			})
		}

		function BindWindowBeforeUnload() {
			$(window).bind("beforeunload", function() {
				if (FormDataWasChanged) return "Some data was changed and not saved. Are you sure want to leave page?"
			})
		}

		function MakeTextAreasAutoSized() {
			autosize($("textarea.autosize"))
		}

		function MakeTextAreasWysiwyg() {
			$("textarea.wysiwyg").each(function() {
				var e = $(this)[0],
					t = e.id,
					n = e.name,
					i = e.innerHTML;
				$(e).wrap('<div id="' + t + '_wysiwyg" class="wysiwyg" style="height:auto;padding:0px"></div>').remove();
				//$(e).wrap('<div id="' + t + '_wysiwyg" class="form-control wysiwyg" style="height:auto;padding:0px"></div>').remove();
				var o = document.getElementById(t + "_wysiwyg");
				o.id = t;
				var a = window.pell.init({
					element: o,
					onChange: function(t) {
						e.innerHTML = HtmlFixEditorHtml(t), $(e).val(HtmlFixEditorHtml(t)), FormDataWasChanged = FormDataWasChanged || !0
					},
					defaultParagraphSeparator: "p",
					styleWithCSS: !1,
					actions: ["paragraph", "heading1", "heading2", "bold", "italic", "underline", "strikethrough", "ulist", "olist", "link", {
						name: "htmlcode",
						icon: "HTML",
						title: "HTML Source",
						result: function(t, e, n) {
							var i = $(t),
								o = $(e),
								a = $(n);
							if (a.hasClass("pell-button-html-pressed")) {
								a.removeClass("pell-button-html-pressed"), i.removeClass("pell-html-mode"), i.find(".pell-actionbar .pell-button").prop("disabled", !1);
								var s = i.find("textarea.form-control").val();
								e.innerHTML = HtmlFixEditorHtml(s), $(e).val(HtmlFixEditorHtml(s)), setTimeout(function() {
									o.focus()
								}, 0)
							} else a.addClass("pell-button-html-pressed"), i.addClass("pell-html-mode"), i.find(".pell-actionbar .pell-button").prop("disabled", !0), a.prop("disabled", !1), setTimeout(function() {
								i.find("textarea.form-control").focus()
							}, 0)
						}
					}],
					classes: {
						actionbar: "pell-actionbar",
						button: "pell-button",
						content: "pell-content",
						selected: "pell-button-selected"
					}
				});
				a.onfocusin = function() {
					$(o).addClass("focused")
				}, a.onfocusout = function() {
					$(o).find(".pell-actionbar button.pell-button-selected").removeClass("pell-button-selected"), $(o).removeClass("focused")
				}, $(o).append('<textarea class="form-control" id="' + t + '_wysiwyg" name="' + n + '" style="display:none"></textarea>'), e = document.getElementById(t + "_wysiwyg"), $(e).on("input", function() {
					FormDataWasChanged = FormDataWasChanged || !0
				}), e.innerHTML = HtmlFixEditorHtml(HtmlDecode(i)), $(e).val(HtmlFixEditorHtml(HtmlDecode(i))), a.content.innerHTML = HtmlFixEditorHtml(HtmlDecode(i))
			})
		}

		function MakeTextAreasTmplEditor() {
			var i = !0;
			$("textarea.tmpl-editor").each(function() {
				var e = $(this)[0],
					t = $(this).data("emode"),
					n = "text/html";
				"js" == t ? n = "javascript" : "css" == t && (n = "css"), CodeMirror.fromTextArea(e, {
					lineNumbers: !0,
					lineWrapping: !0,
					viewportMargin: 1 / 0,
					indentWithTabs: !0,
					indentUnit: 4,
					tabSize: 4,
					mode: n
				}).on("change", function(t) {
					e.value = t.getValue(), i || (FormDataWasChanged = FormDataWasChanged || !0)
				})
			}), i = !1
		}

		function MakeTextAreasNotReactOnTab() {
			$("textarea.use-tab-key").each(function() {
				$(this).keydown(function(t) {
					if (9 === t.keyCode) {
						var e = this.selectionStart,
							n = this.selectionEnd,
							i = $(this),
							o = i.val();
						i.val(o.substring(0, e) + "\t" + o.substring(n)), this.selectionStart = this.selectionEnd = e + 1, t.preventDefault(), FormDataWasChanged = FormDataWasChanged || !0
					}
				})
			})
		}

		function Initialize() {
			"function" == typeof $ ? (AllFormsToAjax(), BindWindowBeforeUnload(), MakeTextAreasAutoSized(), MakeTextAreasWysiwyg(), MakeTextAreasTmplEditor(), MakeTextAreasNotReactOnTab()) : console.log("Error: jQuery is not loaded!")
		}
		return window.addEventListener ? window.addEventListener("load", Initialize, !1) : window.attachEvent && window.attachEvent("onload", Initialize), {
			ShowMsgSuccess: function(t, e) {
				ShowSystemMsg(t, e, !1)
			},
			ShowMsgError: function(t, e) {
				ShowSystemMsg(t, e, !0)
			},
			FormDataWasChanged: function() {
				PreventDataLost()
			},
			ModalUserProfile: function() {
				var t = '<div class="modal fade" id="sys-modal-user-settings" tabindex="-1" role="dialog" aria-labelledby="sysModalUserSettingsLabel" aria-hidden="true"> \t\t\t\t\t<div class="modal-dialog modal-dialog-centered" role="document"> \t\t\t\t\t\t<div class="modal-content"> \t\t\t\t\t\t\t<form class="form-user-settings" action="/cp/" method="post" autocomplete="off"> \t\t\t\t\t\t\t\t<input type="hidden" name="action" value="index-user-update-profile"> \t\t\t\t\t\t\t\t<div class="modal-header"> \t\t\t\t\t\t\t\t\t<h5 class="modal-title" id="sysModalUserSettingsLabel">Настройки профиля</h5> \t\t\t\t\t\t\t\t\t<button type="button" class="close" data-dismiss="modal" aria-label="Close"> \t\t\t\t\t\t\t\t\t\t<span aria-hidden="true">&times;</span> \t\t\t\t\t\t\t\t\t</button> \t\t\t\t\t\t\t\t</div> \t\t\t\t\t\t\t\t<div class="modal-body text-left"> \t\t\t\t\t\t\t\t\t<div class="form-group"> \t\t\t\t\t\t\t\t\t\t<label for="first_name">Логин</label> \t\t\t\t\t\t\t\t\t\t<input type="text" class="form-control" id="first_name" name="first_name" value="' + window.CurrentUserProfileData.first_name + '" placeholder="Логин пользователя" autocomplete="off"> \t\t\t\t\t\t\t\t\t</div> \t\t\t\t\t\t\t\t\t<div class="form-group"> \t\t\t\t\t\t\t\t\t\t<label for="last_name">Алиас</label> \t\t\t\t\t\t\t\t\t\t<input type="text" class="form-control" id="last_name" name="last_name" value="' + window.CurrentUserProfileData.last_name + '" placeholder="Алиас пользователя" autocomplete="off"> \t\t\t\t\t\t\t\t\t</div> \t\t\t\t\t\t\t\t\t<div class="form-group"> \t\t\t\t\t\t\t\t\t\t<label for="email">Email</label> \t\t\t\t\t\t\t\t\t\t<input type="email" class="form-control" id="email" name="email" value="' + window.CurrentUserProfileData.email + '" placeholder="Email пользователя" autocomplete="off" required> \t\t\t\t\t\t\t\t\t</div> \t\t\t\t\t\t\t\t\t<div class="form-group"> \t\t\t\t\t\t\t\t\t\t<label for="address">IP Адрес сервера приложений</label> \t\t\t\t\t\t\t\t\t\t<input type="text" pattern="^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$" class="form-control" id="address" name="address" value="' + window.CurrentUserProfileData.address + '" placeholder="Адрес сервера приложения" autocomplete="off" required> \t\t\t\t\t\t\t\t\t</div> \t\t\t\t\t\t\t\t\t<div class="form-group"> \t\t\t\t\t\t\t\t\t\t<label for="port">Порт сервера</label> \t\t\t\t\t\t\t\t\t\t<input type="number" min="0" step="1" class="form-control" id="port" name="port" value="' + window.CurrentUserProfileData.port + '" placeholder="Порт сервера" autocomplete="off" required> \t\t\t\t\t\t\t\t\t</div> \t\t\t\t\t\t\t\t\t<div class="form-group"> \t\t\t\t\t\t\t\t\t\t<label for="password">New password</label> \t\t\t\t\t\t\t\t\t\t<input type="password" class="form-control" id="password" aria-describedby="passwordHelp" name="password" value="" placeholder="User new password" autocomplete="off"> \t\t\t\t\t\t\t\t\t\t<small id="passwordHelp" class="form-text text-muted">Оставьте это поле пустым, если не хотите изменить пароль</small> \t\t\t\t\t\t\t\t\t</div> \t\t\t\t\t\t\t\t\t<div class="sys-messages"></div> \t\t\t\t\t\t\t\t</div> \t\t\t\t\t\t\t\t<div class="modal-footer"> \t\t\t\t\t\t\t\t\t<button type="submit" class="btn btn-primary">Save</button> \t\t\t\t\t\t\t\t\t<button type="button" class="btn btn-secondary" data-dismiss="modal">Cancel</button> \t\t\t\t\t\t\t\t</div> \t\t\t\t\t\t\t</form> \t\t\t\t\t\t</div> \t\t\t\t\t</div> \t\t\t\t</div>';
				$("#sys-modal-user-settings-placeholder").html(t), $("#sys-modal-user-settings").modal({
					backdrop: "static",
					keyboard: !1,
					show: !1
				}), $("#sys-modal-user-settings").on("hidden.bs.modal", function(t) {
					$("#sys-modal-user-settings-placeholder").html("")
				}), FormToAjax($("#sys-modal-user-settings form")), $("#sys-modal-user-settings").modal("show")
			},
			ShopProductsAdd: function() {
				var selText = $("#lbl_attributes option:selected").text(),
					selValue = $("#lbl_attributes").val();
				"0" != selValue && ($("#lbl_attributes")[0].selectedIndex = 0, $("#lbl_attributes").selectpicker("refresh"), 0 < $("#prod_attr_" + selValue).length || ($("#list").append('<div class="form-group" id="prod_attr_' + selValue + '"><div><b>' + selText + '</b></div><div class="position-relative"><select class="form-control" name="value.' + selValue + '" autocomplete="off" required disabled><option value="0">Loading values...</option></select><button type="button" class="btn btn-danger btn-dynamic-remove" onclick="fave.ShopProductsRemove(this);" disabled>&times;</button></div></div>'), PreventDataLost(), $.ajax({
					type: "POST",
					url: "/cp/",
					data: {
						action: "shop-get-attribute-values",
						id: selValue
					}
				}).done(function(data) {
					try {
						eval(data)
					} catch (t) {
						t instanceof SyntaxError && (console.log(data), console.log("Error: JavaScript code eval error", t.message))
					}
				}).fail(function(xhr, status, error) {
					$("#prod_attr_" + selValue).remove();
					try {
						eval(xhr.responseText)
					} catch (t) {
						t instanceof SyntaxError && (console.log(xhr.responseText), console.log("Error: JavaScript code eval error", t.message))
					}
				})))
			},
			ShopProductsRemove: function(t) {
				$(t).parent().parent().remove(), PreventDataLost()
			},
			ShopAttributesAdd: function() {
				$("#list").append('<div class="form-group position-relative"><input class="form-control" type="text" name="value.0" value="" placeholder="" autocomplete="off" required><button type="button" class="btn btn-danger btn-dynamic-remove" onclick="fave.ShopAttributesRemove(this);">&times;</button></div>'), PreventDataLost(), setTimeout(function() {
					$("#list input").last().focus()
				}, 100)
			},
			ShopAttributesRemove: function(t) {
				$(t).parent().remove(), PreventDataLost()
			},
			ShopProductsUploadImage: function(action_name, product_id, input_id) {
				var file_el = $("#" + input_id)[0];
				if (file_el.files && !(file_el.files.length <= 0)) {
					$("#img-upload-block input").prop("disabled", !0), $("#upload-msg").css("display", "block");
					var fd = new FormData;
					fd.append("action", action_name), fd.append("id", product_id), fd.append("count", file_el.files.length);
					for (var i = 0; i < file_el.files.length; i++) fd.append("file_" + i, file_el.files[i]);
					$.ajax({
						url: "/cp/",
						method: "POST",
						type: "POST",
						data: fd,
						contentType: !1,
						processData: !1
					}).done(function(data) {
						try {
							eval(data)
						} catch (t) {
							t instanceof SyntaxError && (console.log(data), console.log("Error: JavaScript code eval error", t.message))
						}
					}).fail(function(xhr, status, error) {
						try {
							eval(xhr.responseText)
						} catch (t) {
							t instanceof SyntaxError && (console.log(xhr.responseText), console.log("Error: JavaScript code eval error", t.message))
						}
					}).always(function() {
						file_el.value = "", $("#img-upload-block input").prop("disabled", !1), $("#upload-msg").css("display", "none")
					})
				}
			},
			ShopProductsDeleteImage: function(t, e, n) {
				$(t).hasClass("in-progress") || ($(t).addClass("in-progress"), $.ajax({
					type: "POST",
					url: "/cp/",
					data: {
						action: "shop-upload-delete",
						id: e,
						file: n
					}
				}).done(function(t) {
					IsDebugMode() && console.log("done", t), AjaxDone(t)
				}).fail(function(t, e, n) {
					IsDebugMode() && console.log("fail", t, e, n), AjaxFail(t.responseText, e, n)
				}))
			},
			ShopProductsDuplicateBase: function(button, product_id, attach) {
				$(button).hasClass("in-progress") || (FormDataIsChanged() ? fave.ShowMsgError("Warning!", "Something was changed, save changes before duplicate product", !0) : ($(button).addClass("in-progress"), $.ajax({
					type: "POST",
					url: "/cp/",
					data: {
						action: "shop-duplicate",
						id: product_id,
						attach: attach
					}
				}).done(function(data) {
					try {
						eval(data)
					} catch (t) {
						t instanceof SyntaxError && (console.log(data), console.log("Error: JavaScript code eval error", t.message))
					}
				}).fail(function(xhr, status, error) {
					try {
						eval(xhr.responseText)
					} catch (t) {
						t instanceof SyntaxError && (console.log(xhr.responseText), console.log("Error: JavaScript code eval error", t.message))
					}
				}).always(function() {
					$(button).removeClass("in-progress")
				})))
			},
			ShopProductsDuplicate: function(t, e) {
				fave.ShopProductsDuplicateBase(t, e, 0)
			},
			ShopProductsDuplicateWithAttach: function(t, e) {
				fave.ShopProductsDuplicateBase(t, e, 1)
			},
			ShopProductsRetryImage: function(t, e) {
				var n = $("#" + e),
					i = n.attr("src");
				n.attr("src", "/assets/cp/img-load.gif"), setTimeout(function() {
					n.attr("src", i)
				}, 1e3)
			},
			ActionLogout: function(t, url_ = "/cp/") {
				confirm(t) && $.ajax({
					type: "POST",
					url: url_,
					data: {
						pier: url_,
						action: "index-user-logout"
					}
				}).done(function(t) {
					IsDebugMode() && console.log("done", t), AjaxDone(t)
				}).fail(function(t, e, n) {
					IsDebugMode() && console.log("fail", t, e, n), AjaxFail(t.responseText, e, n)
				})
			},
			ActionDataTableDelete: function(t, e, n, i) {
				confirm(i) && $.ajax({
					type: "POST",
					url: "/cp/",
					data: {
						action: e,
						id: n
					}
				}).done(function(t) {
					IsDebugMode() && console.log("done", t), AjaxDone(t)
				}).fail(function(t, e, n) {
					IsDebugMode() && console.log("fail", t, e, n), AjaxFail(t.responseText, e, n)
				})
			},
			ActionThemeFile: function(t, e, n) {
				return confirm(n) && $.ajax({
					type: "POST",
					url: "/cp/",
					data: {
						action: t,
						file: e
					}
				}).done(function(t) {
					IsDebugMode() && console.log("done", t), AjaxDone(t)
				}).fail(function(t, e, n) {
					IsDebugMode() && console.log("fail", t, e, n), AjaxFail(t.responseText, e, n)
				}), !1
			},
			ShopProductsImageReorder: function(t, e) {
				$.ajax({
					type: "POST",
					url: "/cp/",
					data: {
						action: t,
						data: JSON.stringify(e)
					}
				}).done(function(t) {
					IsDebugMode() && console.log("done", t), AjaxDone(t)
				}).fail(function(t, e, n) {
					IsDebugMode() && console.log("fail", t, e, n), AjaxFail(t.responseText, e, n)
				})
			},
			ShopAttachProduct: function(e) {
				$("#sys-modal-shop-product-attach-placeholder").html('<div class="modal fade" id="sys-modal-shop-product-attach" tabindex="-1" role="dialog" aria-labelledby="sysModalShopProductLabel" aria-hidden="true"> \t\t\t\t\t<div class="modal-dialog modal-dialog-centered" role="document"> \t\t\t\t\t\t<div class="modal-content"> \t\t\t\t\t\t\t<div class="modal-header"> \t\t\t\t\t\t\t\t<h5 class="modal-title" id="sysModalShopProductLabel">Attach product</h5> \t\t\t\t\t\t\t\t<button type="button" class="close" data-dismiss="modal" aria-label="Close"> \t\t\t\t\t\t\t\t\t<span aria-hidden="true">&times;</span> \t\t\t\t\t\t\t\t</button> \t\t\t\t\t\t\t</div> \t\t\t\t\t\t\t<div class="modal-body text-left"> \t\t\t\t\t\t\t\t<div class="form-group"> \t\t\t\t\t\t\t\t\t<input type="text" class="form-control" name="product-name" value="" placeholder="Type product name here..." readonly autocomplete="off"> \t\t\t\t\t\t\t\t</div> \t\t\t\t\t\t\t\t<div class="form-group" style="margin-bottom:0px;"> \t\t\t\t\t\t\t\t\t<div class="products-list"></div> \t\t\t\t\t\t\t\t</div> \t\t\t\t\t\t\t</div> \t\t\t\t\t\t\t<div class="modal-footer"> \t\t\t\t\t\t\t\t<button type="button" class="btn btn-secondary" data-dismiss="modal">Cancel</button> \t\t\t\t\t\t\t</div> \t\t\t\t\t\t</div> \t\t\t\t\t</div> \t\t\t\t</div>'), $("#sys-modal-shop-product-attach").modal({
					backdrop: "static",
					keyboard: !0,
					show: !1
				}), $("#sys-modal-shop-product-attach").on("hidden.bs.modal", function(t) {
					$("#sys-modal-shop-product-attach-placeholder").html("")
				}), $("#sys-modal-shop-product-attach").modal("show"), setTimeout(function() {
					var t = $('#sys-modal-shop-product-attach input[name="product-name"]');
					t.keyup(function() {
						$.ajax({
							type: "POST",
							url: "/cp/",
							data: {
								action: "shop-attach-product-search",
								words: this.value,
								id: e
							}
						}).done(function(t) {
							0 < $("#sys-modal-shop-product-attach").length && (IsDebugMode() && console.log("done", t), AjaxDone(t))
						}).fail(function(t, e, n) {
							0
						})
					}), t.attr("readonly", !1), t.keyup(), t.focus()
				}, 500)
			},
			ShopAttachProductTo: function(t, e) {
				$.ajax({
					type: "POST",
					url: "/cp/",
					data: {
						action: "shop-attach-product-to",
						parent_id: t,
						product_id: e
					}
				}).done(function(t) {
					IsDebugMode() && console.log("done", t), AjaxDone(t)
				}).fail(function(t, e, n) {
					IsDebugMode() && console.log("fail", t, e, n), AjaxFail(t.responseText, e, n)
				})
			},
			ShopSetOrderStatus: function(t, e, n, i) {
				confirm(i) && $.ajax({
					type: "POST",
					url: "/cp/",
					data: {
						action: "shop-order-set-status",
						id: e,
						status: n
					}
				}).done(function(t) {
					IsDebugMode() && console.log("done", t), AjaxDone(t)
				}).fail(function(t, e, n) {
					IsDebugMode() && console.log("fail", t, e, n), AjaxFail(t.responseText, e, n)
				})
			},
			FilesManagerDialog: function() {
				$("#sys-modal-files-manager-placeholder").html('<div class="modal fade" id="sys-modal-files-manager" tabindex="-1" role="dialog" aria-labelledby="sysModalFilesManagerLabel" aria-hidden="true"> \t\t\t\t\t<div class="modal-dialog modal-dialog-centered" role="document"> \t\t\t\t\t\t<div class="modal-content"> \t\t\t\t\t\t\t<input type="hidden" name="path" value="/"> \t\t\t\t\t\t\t<div class="modal-header"> \t\t\t\t\t\t\t\t<h5 class="modal-title" id="sysModalFilesManagerLabel">Files manager</h5> \t\t\t\t\t\t\t\t<button type="button" class="close" data-dismiss="modal" aria-label="Close"> \t\t\t\t\t\t\t\t\t<span aria-hidden="true">&times;</span> \t\t\t\t\t\t\t\t</button> \t\t\t\t\t\t\t</div> \t\t\t\t\t\t\t<div class="modal-body text-left"> \t\t\t\t\t\t\t\t<div class="dialog-path alert alert-secondary"><span class="text-dotted">/</span></div> \t\t\t\t\t\t\t\t<div class="dialog-data"></div> \t\t\t\t\t\t\t</div> \t\t\t\t\t\t\t<div class="modal-footer"> \t\t\t\t\t\t\t\t<button type="button" class="btn btn-success upload" disabled>Uploading...</button> \t\t\t\t\t\t\t\t<input class="form-control" type="file" id="fmfiles" name="fmfiles" onchange="fave.FilesManagerUploadFile();" style="font-size:12px;background-color:#28a745;border-color:#28a745;color:#fff;cursor:pointer;" multiple=""> \t\t\t\t\t\t\t\t<button type="button" class="btn btn-primary folder" onclick="fave.FilesManagerNewFolderClick();" disabled>New folder</button> \t\t\t\t\t\t\t\t<button type="button" class="btn btn-secondary" data-dismiss="modal">Close</button> \t\t\t\t\t\t\t</div> \t\t\t\t\t\t</div> \t\t\t\t\t</div> \t\t\t\t</div>'), $("#sys-modal-files-manager").modal({
					backdrop: "static",
					keyboard: !0,
					show: !1
				}), $("#sys-modal-files-manager").on("hidden.bs.modal", function(t) {
					$("#sys-modal-files-manager-placeholder").html("")
				}), $("#sys-modal-files-manager").modal("show"), setTimeout(function() {
					fave.FilesManagerLoadData("/")
				}, 500)
			},
			FilesManagerSetPath: function(t) {
				$("#sys-modal-files-manager input[name=path]").val(t), $("#sys-modal-files-manager .dialog-path span").html(t)
			},
			FilesManagerGetPath: function() {
				return $("#sys-modal-files-manager input[name=path]").val()
			},
			FilesManagerRemoveFolder: function(t, e) {
				confirm(e) && (fave.FilesManagerEnableDisableButtons(!0), $.ajax({
					type: "POST",
					url: "/cp/",
					data: {
						action: "files-remove-folder",
						file: t
					}
				}).done(function(t) {
					0 < $("#sys-modal-files-manager").length && (IsDebugMode() && console.log("done", t), AjaxDone(t))
				}).fail(function(t, e, n) {
					0 < $("#sys-modal-files-manager").length && (IsDebugMode() && console.log("fail", t, e, n), AjaxFail(t.responseText, e, n))
				}).always(function() {
					fave.FilesManagerEnableDisableButtons(!1)
				}))
			},
			FilesManagerRemoveFile: function(t, e) {
				confirm(e) && (fave.FilesManagerEnableDisableButtons(!0), $.ajax({
					type: "POST",
					url: "/cp/",
					data: {
						action: "files-remove-file",
						file: t
					}
				}).done(function(t) {
					0 < $("#sys-modal-files-manager").length && (IsDebugMode() && console.log("done", t), AjaxDone(t))
				}).fail(function(t, e, n) {
					0 < $("#sys-modal-files-manager").length && (IsDebugMode() && console.log("fail", t, e, n), AjaxFail(t.responseText, e, n))
				}).always(function() {
					fave.FilesManagerEnableDisableButtons(!1)
				}))
			},
			FilesManagerLoadData: function(t) {
				fave.FilesManagerEnableDisableButtons(!0), $("#sys-modal-files-manager .dialog-data").html('<div class="fm-loading"></div>'), $.ajax({
					type: "POST",
					url: "/cp/",
					data: {
						action: "files-list",
						path: t
					}
				}).done(function(t) {
					0 < $("#sys-modal-files-manager").length && (IsDebugMode() && console.log("done", t), AjaxDone(t))
				}).fail(function(t, e, n) {
					0 < $("#sys-modal-files-manager").length && (IsDebugMode() && console.log("fail", t, e, n), AjaxFail(t.responseText, e, n))
				})
			},
			FilesManagerLoadDataUp: function(t) {
				newPath = t.replace(/\/$/i, ""), newPath = newPath.replace(/[^\/]+$/i, ""), fave.FilesManagerLoadData(newPath)
			},
			FilesManagerEnableDisableButtons: function(t) {
				$("#sys-modal-files-manager #fmfiles").prop("disabled", t), $("#sys-modal-files-manager button.folder").prop("disabled", t)
			},
			FilesManagerUploadFile: function() {
				var t = $("#fmfiles")[0];
				if (t.files && !(t.files.length <= 0)) {
					fave.FilesManagerEnableDisableButtons(!0), $("#sys-modal-files-manager").addClass("uploading");
					var e = new FormData;
					e.append("action", "files-upload"), e.append("count", t.files.length), e.append("path", fave.FilesManagerGetPath());
					for (var n = 0; n < t.files.length; n++) e.append("file_" + n, t.files[n]);
					$.ajax({
						url: "/cp/",
						method: "POST",
						type: "POST",
						data: e,
						contentType: !1,
						processData: !1
					}).done(function(t) {
						0 < $("#sys-modal-files-manager").length && (IsDebugMode() && console.log("done", t), AjaxDone(t))
					}).fail(function(t, e, n) {
						0 < $("#sys-modal-files-manager").length && (IsDebugMode() && console.log("fail", t, e, n), AjaxFail(t.responseText, e, n))
					}).always(function() {
						t.value = "", $("#sys-modal-files-manager").removeClass("uploading"), fave.FilesManagerEnableDisableButtons(!1)
					})
				}
			},
			FilesManagerNewFolderClick: function() {
				var t = prompt("Please enter new folder name", "");
				null != t && (path = fave.FilesManagerGetPath(), fave.FilesManagerEnableDisableButtons(!0), $.ajax({
					type: "POST",
					url: "/cp/",
					data: {
						action: "files-mkdir",
						path: path,
						name: t
					}
				}).done(function(t) {
					0 < $("#sys-modal-files-manager").length && (IsDebugMode() && console.log("done", t), AjaxDone(t))
				}).fail(function(t, e, n) {
					0 < $("#sys-modal-files-manager").length && (IsDebugMode() && console.log("fail", t, e, n), AjaxFail(t.responseText, e, n))
				}).always(function() {
					fave.FilesManagerEnableDisableButtons(!1)
				}))
			}
		}
	}(window, $);
	window.fave = fave
}(window, jQuery);`)
