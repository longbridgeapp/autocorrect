0.8.0
--------

- Add `Unformat` / `UnformatHTML` method for remove spacings.

0.7.0
--------

- Use new HTML parser to format HTML for performance up.
- Avoid format text with script/style/textarea/pre tags.

0.6.1
--------

- Fix Halfwidth to correct fullwidth spaces.

0.6.0
--------

- Auto correct FullWidth -> HalfWidth for Letters, Numbers, and Colon in time.

0.4.1
--------

- Avoid create regex on format method call for performance up (~40%).

0.4.0
--------

- Add Full CJK (Chinese, Japanese, Korean) support.

0.3.3
--------

- Fix space around `-`;

0.3.2
--------

- Fix add space round `*`;

0.3.1
--------

- Fix HTML replace when content has escapeable `&amp;`, `&nbsp;` chars.

0.3.0
--------

- Rename package from `autospace` to `autocorrect`.

0.2.0
--------

- Add `FormatHTML` method for process HTML contents.

0.1.0
--------

- First release.
