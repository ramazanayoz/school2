use strict;
use warnings;

my $file = $ARGV[0];

open IN, '<', $file or die "$file not found!: $!\n";

my @lines = <IN>;
close IN;

my %word_counts = ();

foreach my $line (@lines) {
    chomp $line;
    my @row = split(" ", $line);
    foreach my $word (@row){
        if (exists $word_counts{$word}){
          $word_counts{$word}++;
        }else {
          $word_counts{$word} = 1;
        }
    }
}

my @words = keys %word_counts;
foreach my $w (@words) {
  print "$w -> $word_counts{$w}\n";
}
