# text-pic

This tool generates a social media-friendly graphic from some text. It's especially made to be paired with [WriteFreely](https://writefreely.org), a minimalist blogging platform.

## Usage

To run this, you'll need Go installed. With that, you can compile the project:

```
go get github.com/writeas/text-pic/cmd/wfgraphic-cli
```

Then you can run `wfgraphic-cli` with the options below.

```
Usage of wfgraphic-cli:
  -font string
        Post font (options: "serif", "sans", "mono") (default "serif")
  -h string
        WriteFreely instance hostname (e.g. "pencil.writefree.ly")
  -o string
        Image output filename (default "out.png")
  -size string
        Image size, either a single number for a square (e.g. "900") or a combined width and height (e.g. "1080x1920") (default "1024")
  -u string
        WriteFreely author username (for multi-user instances)
```

The actual content of the graphic is read from `stdin`, which you can either supply when prompted to, or pipe in, for example:

```
cat quote.txt | ./wfgraphic-cli -h write.as -u blog
```

## Examples

<table>
<tr>
<td>
    <img src="https://i.snap.as/NHrWle49.png" alt="Sample using default fonts." />
</td>
<td>
    <img src="https://i.snap.as/8SI7lfyb.png" alt="Sample using 'sans' font and a custom domain." />
</td>
</tr>
</table>