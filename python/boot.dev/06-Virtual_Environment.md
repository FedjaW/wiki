# Virtual Environment (venv)

Virtual environments are Python's way to keep dependencies (e.g. the pygame module) separate from other projects on our machine. For example, we need pygame version 2 for this project, but another project on your computer might require version 1.

As a best practice, each Python project on your machine should have its own virtual environment to keep them isolated from each other.

Create a virtual environment at the top level of your project directory:

```cli
python3 -m venv venv
```

Activate the virtual environment:

```cli
source venv/bin/activate
```

In a requirements.txt the dependencies are defined: pygame==2.6.0

Install the requirements:

```cli
pip install -r requirements.txt
```

pip is Python's package manager. It will install the pygame module into the virtual environment you created.