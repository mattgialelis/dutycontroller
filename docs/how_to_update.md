

## How to Update thesse docs

For more detailed information on how to contribute to this project, please refer to the Contributing file.


### Local Development

Install Material for MkDocs with `pip`, ideally by using a [virtual environment](https://realpython.com/what-is-pip/#using-pip-in-a-python-virtual-environment).

```bash
pip install mkdocs-material mkdocs-git-revision-date-localized-plugin mkdocs-awesome-pages-plugin
```

In the repository root run:
```bash
$ mkdocs serve
INFO     -  Building documentation...
INFO     -  Cleaning site directory
INFO     -  Documentation built in 0.58 seconds
INFO     -  [09:52:56] Watching paths for changes: 'docs', 'mkdocs.yaml'
INFO     -  [09:52:56] Serving on http://127.0.0.1:8000/
INFO     -  [09:52:57] Browser connected: http://localhost:8000/
```
Edit markdown files in the `docs/` folder or the `mkdocs.yaml` file.
Visit [http://localhost:8000/](http://localhost:8000/) to see your docs rendered live as you type.

Once done, please create PR.
